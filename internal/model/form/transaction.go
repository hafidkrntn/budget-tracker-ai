package form

import (
	"time"

	"github.com/google/uuid"
)

type TransactionForm struct {
	ID         uuid.UUID `json:"id"`
	Type       string    `json:"type" binding:"required,oneof=income expense"`
	Amount     float64   `json:"amount" binding:"required,gt=0"`
	Date       time.Time `json:"date" binding:"required"`
	Note       string    `json:"note"`
	UserId     uuid.UUID `json:"user_id" binding:"required"` // harus UUID valid
	CategoryId uuid.UUID `json:"category_id" binding:"required"`
}

type TransactionParams struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}
