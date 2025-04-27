package validator

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

func NewFielValidation(field, msg string) FieldError {
	return FieldError{
		Field:   field,
		Message: msg,
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
