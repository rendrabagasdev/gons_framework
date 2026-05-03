package middlewares

import "github.com/gofiber/fiber/v3"

func AuthGuard() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		token := ctx.Get("Authorization")
		if token == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Accees Denied",
			})
		}

		// Free Logic Here

		return ctx.Next()
	}
}
