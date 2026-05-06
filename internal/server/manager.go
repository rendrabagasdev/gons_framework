package server

import (
	"gons/pkg/env"
	"gons/internal/middleware"
	"gons/internal/vite"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v3"
	"strings"
)

// NewServer initializes and returns a new Fiber application.
func NewServer() *fiber.App {
	engine := html.New("./resources/views", ".html")
	engine.AddFunc("vite", vite.ViteHelper)

	app := fiber.New(fiber.Config{
		AppName:      env.Get("APP_NAME", "Gons Framework"),
		Views:        engine,
		ErrorHandler: middleware.ErrorHandler,
	})

	app.Hooks().OnPreStartupMessage(func(sm *fiber.PreStartupMessageData) error {
		sm.PreventDefault = true
		return nil
	})
	
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} | ${latency} | ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   env.Get("APP_TIMEZONE", "Asia/Jakarta"),
	}))

	allowOrigins := strings.Split(env.Get("CORS_ALLOW_ORIGINS", "*"), ",")
	allowHeaders := strings.Split(env.Get("CORS_ALLOW_HEADERS", "Origin,Content-Type,Accept,Authorization"), ",")

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowHeaders: allowHeaders,
	}))

	app.Use("/", static.New("./public"))

	return app
}
