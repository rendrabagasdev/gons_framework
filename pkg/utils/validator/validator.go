package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var Validate = validator.New()
func ValidatorRequest(c fiber.Ctx, payload any) map[string]string {
	if err := c.Bind().Body(payload); err != nil {
		return map[string]string{"error": "Invalid JSON request format"}
	}

	err := Validate.Struct(payload)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return map[string]string{"error": "Internal error: Invalid payload structure passed to validator"}
		}

		validationErrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return map[string]string{"error": "Unknown validation error"}
		}

		errors := make(map[string]string)
		for _, e := range validationErrs {
			errors[e.Field()] = fmt.Sprintf("The %s field is %s", e.Field(), e.Tag())
		}

		return errors
	}

	return nil
}
