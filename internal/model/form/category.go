package form

import "github.com/google/uuid"

type Category struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type CategoryParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
