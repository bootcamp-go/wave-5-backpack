package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GenerateMessageValidationError(errors validator.ValidationErrors) []string {
	invalidErrors := make([]string, 0)
	for _, e := range errors {
		field := e.Field()
		invalidErrors = append(invalidErrors, fmt.Sprintf("el campo %s es requerido",
			field))
	}
	return invalidErrors
}

func GenerateSimpleErrorMessage(err error) string {
	return err.Error()
}
