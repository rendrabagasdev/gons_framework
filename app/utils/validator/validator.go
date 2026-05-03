package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var Validate = validator.New()

func ValidatorRequest(c fiber.Ctx, payload any) map[string]string {
	if err := c.Bind().Body(payload); err != nil {
		return map[string]string{"error": "require JSON request"}
	}

	err := Validate.Struct(payload)
	if err != nil {
		errors := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = fmt.Sprintf("%s field is %s", err.Field(), err.Tag())
		}

		return errors
	}

	return nil
}
