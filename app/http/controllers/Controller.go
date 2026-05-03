package controllers

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	"github.com/golobby/container/v3"
)

type Controller struct{}

func NewController() *Controller {
	controller := &Controller{}
	if err := container.Fill(controller); err != nil {
		slog.Error("Gons: Controller injection error", "error", err)
	}
	return controller
}

func (controller *Controller) Welcome(c fiber.Ctx) error {
	return c.Render("welcome", fiber.Map{
		"Title": "Welcome",
	}, "layouts/main")
}
