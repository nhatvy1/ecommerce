package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

func ValidateFunc[T any](req T) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(req)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			errors = append(errors, ErrorResponse{
				Field:   err.Field(),
				Message: CustomErrMsg(err),
			})
		}
	}
	return errors
}

func CustomErrMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "email":
		return "Invalid email format"
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", err.Field(), err.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("%s is invalid", err.Field())
	}
}
