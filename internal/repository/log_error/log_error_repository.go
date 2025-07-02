package logerror

import (
	"backend-go/internal/model/migrate"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type LogErrorRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogErrorRepository {
	return &LogErrorRepository{db: db}
}

func (r *LogErrorRepository) CreateLogError(msg, file string, code, line int) (int64, error) {
	insertLogError := migrate.LogError{
		ErrorMessage:    msg,
		ErrorFileName:   file,
		ErrorCode:       strconv.Itoa(code),
		ErrorNumberLine: line,
		CreatedAt:       time.Now(),
	}

	if err := r.db.Create(&insertLogError).Error; err != nil {
		return 0, err
	}

	return 1, nil
}
