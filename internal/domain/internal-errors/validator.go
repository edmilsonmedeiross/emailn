package internalerrors

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(input interface{}) error {
	validate := validator.New()
	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	validationError := validationErrors[0]

	var field = strings.ToLower(validationError.Field())

	switch validationError.Tag() {
	case "required":
		return fmt.Errorf("the field %s is required", field)
	case "email":
		return fmt.Errorf("the field %s must be a valid email address", field)
	case "min":
		return fmt.Errorf("the field %s must have a minimum value of %s", field, validationError.Param())
	case "max":
		return fmt.Errorf("the field %s must have a maximum value of %s", field, validationError.Param())
	}

	return nil
}
