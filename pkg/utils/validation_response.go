package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

var validationMessages = map[string]string{
	"required": "{field} is required",
	"min":      "{field} must be at least {param} characters",
	"max":      "{field} must be less than {param} characters",
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Errors  []FieldError `json:"errors"`
}

func MapValidationErrors(err error) []FieldError {
	var errors []FieldError

	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		return errors
	}

	for _, e := range ve {
		field := strings.ToLower(e.Field())
		tag := e.Tag()
		param := e.Param()

		template, exists := validationMessages[tag]
		msg := "is invalid"

		if exists {
			msg = template
			msg = strings.ReplaceAll(msg, "{field}", field)
			msg = strings.ReplaceAll(msg, "{param}", param)
		}

		errors = append(errors, FieldError{
			Field:   field,
			Message: msg,
		})
	}

	return errors
}

func ValidationErrorResponse(
	w http.ResponseWriter,
	err error,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	return json.NewEncoder(w).Encode(ValidationError{
		Status:  false,
		Message: "validation error",
		Errors:  MapValidationErrors(err),
	})
}
