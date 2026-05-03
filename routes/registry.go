package routes

import "github.com/gofiber/fiber/v3"

func RegisterRoute(app *fiber.App) {
	WebRoutes(app)
}
