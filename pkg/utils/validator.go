package utils

import (
	"github.com/go-playground/validator/v10"
)

// NewValidator creates a new validator instance with custom error messages.
func NewValidator() *validator.Validate {
	v := validator.New()

	// Register custom validation messages
	v.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		return true // Example override, replace with your logic
	})

	return v
}

// TranslateError translates validation errors into user-friendly messages.
func TranslateError(err error) string {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			switch e.Tag() {
			case "required":
				return e.Field() + " is required"
			case "email":
				return "Invalid email format"
			default:
				return "Invalid input"
			}
		}
	}
	return "Unknown error"
}
