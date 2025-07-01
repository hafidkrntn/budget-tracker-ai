package migrate

import (
	"time"

	"github.com/google/uuid"
)

type LogError struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ErrorMessage    string
	ErrorCode       string
	ErrorNumberLine int
	ErrorFileName   string
	CreatedAt       time.Time
}
