package middlewares

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator for validating request payloads
type CustomValidator struct {
	validator *validator.Validate
}

// Validate method for Echo
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// NewValidator initializes a new validator
func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
