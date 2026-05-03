package controllers

import (
	"go-framework/app/http/requests"
	"go-framework/app/http/services"
	utils "go-framework/app/utils/validator"
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/golobby/container/v3"
)

var validate = validator.New()

type UserController struct {
	UserService *services.UserService `container:"type"`
}

func NewUserController() *UserController {
	controller := &UserController{}
	if err := container.Fill(controller); err != nil {
		slog.Error("Gons: UserController injection error", "error", err)
	}
	return controller
}

func (controller *UserController) GetAll(c fiber.Ctx) error {
	users, err := controller.UserService.GetAllUsers()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal Mengambil data user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "OK",
		"message": "Berhasil Mengambil data user",
		"data":    users,
	})
}

func (u *UserController) Store(c fiber.Ctx) error {
	var req requests.CreateUserRequest

	if errors := utils.ValidatorRequest(c, &req); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Data tidak valid!",
			"errors":  errors,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"data":    req,
	})
}
