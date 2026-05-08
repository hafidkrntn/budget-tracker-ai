package middleware

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDBWithRetry mendapatkan DB connection dengan retry mechanism
func GetDBWithRetry(ctx context.Context, maxRetries int, db *gorm.DB) (*gorm.DB, error) {
	if db == nil {
		return nil, errors.New("database not initialized")
	}

	var err error
	for i := 0; i < maxRetries; i++ {
		// Cek apakah context sudah cancelled
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		sqlDB, dbErr := db.DB()
		if dbErr != nil {
			err = dbErr
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}

		// Ping untuk memastikan koneksi masih hidup
		pingErr := sqlDB.PingContext(ctx)
		if pingErr != nil {
			err = pingErr
			log.Printf("⚠️ Database ping failed (attempt %d/%d): %v", i+1, maxRetries, pingErr)
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}

		return db, nil
	}

	return nil, errors.New("failed to get database connection: " + err.Error())
}

// ExecuteWithRetry menjalankan fungsi database dengan retry mechanism
func ExecuteWithRetry(ctx context.Context, fn func(*gorm.DB) error, maxRetries int, db *gorm.DB) error {
	var err error

	for i := 0; i < maxRetries; i++ {
		// Cek context
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		database, dbErr := GetDBWithRetry(ctx, 3, db)
		if dbErr != nil {
			err = dbErr
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}

		err = fn(database)
		if err == nil {
			return nil
		}

		// Cek apakah error karena koneksi database
		if isConnectionError(err) {
			waitTime := time.Duration(i+1) * time.Second
			log.Printf("⚠️ Database connection error (attempt %d/%d): %v. Retrying in %v...",
				i+1, maxRetries, err, waitTime)
			time.Sleep(waitTime)
			continue
		}

		// Kalau bukan connection error, langsung return
		return err
	}

	return errors.New("max retries exceeded: " + err.Error())
}

func TunePostgresPool(sqlDB *sql.DB) {
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
}

// startHealthCheck melakukan pengecekan kesehatan koneksi database secara berkala
func StartHealthCheck(ctx context.Context, dsnPg string, dbLogger logger.Interface, db *gorm.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	consecutiveFailures := 0
	maxConsecutiveFailures := 3

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if db == nil {
				continue
			}

			sqlDB, err := db.DB()
			if err != nil {
				log.Printf("⚠️ Health check: failed to get sql.DB: %v", err)
				consecutiveFailures++
			} else {
				pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
				err = sqlDB.PingContext(pingCtx)
				cancel()

				if err != nil {
					consecutiveFailures++
					log.Printf("⚠️ Health check: database ping failed (%d/%d): %v",
						consecutiveFailures, maxConsecutiveFailures, err)

					// Kalau gagal berturut-turut, coba reconnect
					if consecutiveFailures >= maxConsecutiveFailures {
						log.Println("🔄 Attempting to reconnect to database...")
						if ReconnectDB(dsnPg, dbLogger, db) {
							consecutiveFailures = 0
							log.Println("✅ Database reconnected successfully")
						} else {
							log.Println("❌ Failed to reconnect to database")
						}
					}
				} else {
					// Reset counter kalau berhasil
					if consecutiveFailures > 0 {
						log.Println("✅ Health check: database connection restored")
						consecutiveFailures = 0
					}
				}
			}
		}
	}
}

// reconnectDB mencoba reconnect ke database
func ReconnectDB(dsnPg string, dbLogger logger.Interface, db *gorm.DB) bool {
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		gormDB, err := gorm.Open(postgres.Open(dsnPg), &gorm.Config{
			Logger: dbLogger,
		})

		if err == nil {
			sqlDB, dbErr := gormDB.DB()
			if dbErr == nil {
				pingErr := sqlDB.Ping()
				if pingErr == nil {
					TunePostgresPool(sqlDB)
					db = gormDB
					return true
				}
			}
		}

		waitTime := time.Duration(i+1) * 2 * time.Second
		log.Printf("⚠️ Reconnect attempt %d/%d failed. Retrying in %v...", i+1, maxRetries, waitTime)
		time.Sleep(waitTime)
	}

	return false
}

// isConnectionError mengecek apakah error adalah connection error
func isConnectionError(err error) bool {
	if err == nil {
		return false
	}

	// Cek error types yang umum untuk connection issues
	return errors.Is(err, gorm.ErrInvalidDB) ||
		errors.Is(err, sql.ErrConnDone) ||
		errors.Is(err, context.DeadlineExceeded) ||
		containsConnectionErrorKeywords(err.Error())
}

// containsConnectionErrorKeywords mengecek keywords yang umum di connection errors
func containsConnectionErrorKeywords(errMsg string) bool {
	keywords := []string{
		"connection refused",
		"connection reset",
		"broken pipe",
		"no such host",
		"network is unreachable",
		"connection timed out",
		"EOF",
		"driver: bad connection",
	}

	errMsgLower := strings.ToLower(errMsg)
	for _, keyword := range keywords {
		if strings.Contains(errMsgLower, keyword) {
			return true
		}
	}
	return false
}
