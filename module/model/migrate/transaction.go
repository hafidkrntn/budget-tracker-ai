package migrate

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Type       string         `gorm:"type:varchar(50);not null" json:"type"` // e.g., "income", "expense"
	Amount     float64        `gorm:"not null" json:"amount"`
	Date       time.Time      `gorm:"not null" json:"date"`
	Note       string         `gorm:"type:text" json:"note"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	CategoryID uuid.UUID      `gorm:"type:uuid;not null;index" json:"category_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`

	// Relasi
	User     User     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"category"`
}
