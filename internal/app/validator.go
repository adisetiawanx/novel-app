package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if ok := errors.As(err, &validationErrors); ok {
			var errorMessages []string
			for _, vErr := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", vErr.Field(), vErr.Tag()))
			}
			return fmt.Errorf(strings.Join(errorMessages, ", "))
		}
		return err
	}
	return nil
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
