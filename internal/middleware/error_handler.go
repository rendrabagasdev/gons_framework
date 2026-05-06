package middleware

import (
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v3"
)

// ErrorHandler is a Laravel-style global error handler for Gons Framework.
func ErrorHandler(c fiber.Ctx, err error) error {
	// Default status code
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Log error with context
	slog.Error("Unhandled Error",
		"method", c.Method(),
		"path", c.Path(),
		"status", code,
		"error", err.Error(),
	)

	// Check if request expects JSON or is from Inertia
	accept := c.Get("Accept")
	isInertia := c.Get("X-Inertia") != ""

	if strings.Contains(accept, "application/json") || isInertia {
		return c.Status(code).JSON(fiber.Map{
			"message": err.Error(),
			"status":  code,
		})
	}

	// Determine view name
	view := "errors/500"
	if code == fiber.StatusNotFound {
		view = "errors/404"
	}

	// Render HTML view
	if errRender := c.Status(code).Render(view, fiber.Map{
		"message": err.Error(),
		"status":  code,
	}); errRender != nil {
		// Fallback to string if render fails
		return c.Status(code).SendString(err.Error())
	}

	return nil
}
