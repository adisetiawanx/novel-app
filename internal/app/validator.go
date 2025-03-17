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
		if errors.As(err, &validationErrors) {
			var errorMessages []string
			for _, vErr := range validationErrors {
				fieldName := vErr.Field()
				switch vErr.Tag() {
				case "eqfield":
					errorMessages = append(errorMessages, fmt.Sprintf("%s is wrong", fieldName))
				default:
					errorMessages = append(errorMessages, fmt.Sprintf("%s is %s", fieldName, vErr.Tag()))
				}
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
