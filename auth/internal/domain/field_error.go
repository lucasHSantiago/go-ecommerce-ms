package domain

import (
	"fmt"
	"strings"
)

type FieldError struct {
	Field   string
	Message string
}

func (fe FieldError) Error() string {
	return fmt.Sprintf("%s: %s", fe.Field, fe.Message)
}

func NewFieldValidation(field string, err error) FieldError {
	return FieldError{
		Field:   field,
		Message: err.Error(),
	}
}

type ValidationErrors []FieldError

func (v ValidationErrors) Error() string {
	var errors []string
	for _, fe := range v {
		errors = append(errors, fe.Error())
	}
	return strings.Join(errors, "; ")
}

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}

	return nil
}
