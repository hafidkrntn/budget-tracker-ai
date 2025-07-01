package validate

import "github.com/go-playground/validator/v10"

func Required(fl validator.FieldLevel) bool {
	bindingField := fl.Field().String()

	return len(bindingField) > 0
}
