package config

import (
	"backend-go/middleware"
	"context"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbConn *gorm.DB

func InitDB() *gorm.DB {
	// Logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	loadEnv()

	dsnPg := os.Getenv("DB_DSN")
	if dsnPg == "" {
		log.Fatal("❌ DB_DSN environment variable is empty")
	}

	// Retry connection dengan backoff
	var gormDB *gorm.DB
	var err error
	maxRetries := 5

	for i := 0; i < maxRetries; i++ {
		gormDB, err = gorm.Open(postgres.Open(dsnPg), &gorm.Config{
			Logger: newLogger,
		})

		if err == nil {
			sqlDB, dbErr := gormDB.DB()
			if dbErr == nil {
				// Test koneksi dengan ping
				pingErr := sqlDB.Ping()
				if pingErr == nil {
					// Connection pool tuning
					middleware.TunePostgresPool(sqlDB)
					DbConn = gormDB
					log.Println("✅ PostgreSQL connected successfully")

					// Start health check goroutine
					go middleware.StartHealthCheck(context.Background(), dsnPg, newLogger, DbConn)
					return DbConn
				}
				err = pingErr
			} else {
				err = dbErr
			}
		}

		waitTime := time.Duration(i+1) * 2 * time.Second
		log.Printf("⚠️ Failed to connect to PostgreSQL (attempt %d/%d): %v. Retrying in %v...",
			i+1, maxRetries, err, waitTime)
		time.Sleep(waitTime)
	}

	log.Fatal("❌ failed to connect to PostgreSQL after retries:", err)
	return nil
}

func GetDB() *gorm.DB {
	if DbConn == nil {
		log.Fatal("❌ database not initialized")
	}
	return DbConn
}
