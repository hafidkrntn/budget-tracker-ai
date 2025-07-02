package response

import (
	"time"

	"github.com/google/uuid"
)

type TransactionPagination struct {
	ID                  uuid.UUID `json:"id"`
	Type                string    `json:"type"`
	Amount              float64   `json:"amount"`
	Date                time.Time `json:"date"`
	Note                string    `json:"note"`
	UserName            string    `json:"user_name"`
	UserEmail           string    `json:"user_email"`
	CategoryName        string    `json:"category_name"`
	CategoryDescription string    `json:"category_description"`
}
