package routes

import (
	"gons/app/http/controllers"
	"gons/app/http/middlewares"

	"github.com/gofiber/fiber/v3"
)

func WebRoutes(app *fiber.App) {
	ctrl := controllers.NewController()
	userCtrl := controllers.NewUserController()

	// Rute Homepage
	app.Get("/", ctrl.Welcome)

	app.Get("/users", userCtrl.GetAll)
	// OR
	protected := app.Group("/protected")
	protected.Use(middlewares.AuthGuard())
	protected.Get("/users", userCtrl.GetAll)

	// Contoh rute lain
	app.Get("/about", func(c fiber.Ctx) error {
		return c.SendString("Ini halaman About dari Gons!")
	})

}
