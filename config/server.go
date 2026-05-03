package config

import (
	"fmt"
	"log"
	"log/slog"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v3"
	"github.com/golobby/container/v3"
)

func RegisterServer() {
	err := container.Singleton(func() *fiber.App {

		engine := html.New("./resources/views", ".html")
		engine.AddFunc("vite", Vitehelper)

		app := fiber.New(fiber.Config{
			AppName: GetEnv("APP_NAME", "Gons Framework"),
			Views:   engine,
		})

		app.Hooks().OnPreStartupMessage(func(sm *fiber.PreStartupMessageData) error {
			sm.PreventDefault = true
			return nil
		})
		app.Use(logger.New(logger.Config{
			Format:     "[${time}] ${status} | ${latency} | ${method} ${path}\n",
			TimeFormat: "15:04:05",
			TimeZone:   GetEnv("APP_TIMEZONE", "Asia/Jakarta"),
		}))

		allowOrigins := strings.Split(GetEnv("CORS_ALLOW_ORIGINS", "*"), ",")
		allowHeaders := strings.Split(GetEnv("CORS_ALLOW_HEADERS", "Origin,Content-Type,Accept,Authorization"), ",")

		app.Use(cors.New(cors.Config{
			AllowOrigins: allowOrigins,
			AllowHeaders: allowHeaders,
		}))

		app.Use("/", static.New("./public"))

		return app
	})

	if err != nil {
		slog.Error("error when register server", "error", err)
		log.Fatalf("failed to register server %v", err)
	}
}

func init() {
	RegisterConfig(func() error {
		RegisterServer()
		return nil
	})
}

// PrintBanner mencetak ASCII art Gons ke terminal
func PrintBanner() {
	appName := GetEnv("APP_NAME", "Gons Framework")
	address := GetServerAddress() // Memakai fungsi yang kita buat sebelumnya

	customBanner := `
   ______                 
  / ____/___  ____  _____ 
 / / __/ __ \/ __ \/ ___/ 
/ /_/ / /_/ / / / (__  )  
\____/\____/_/ /_/____/   
                          
🚀 %s v1.0.0
--------------------------------------------------
INFO  Server Address:    http://%s
INFO  Environment:       Development
INFO  Routes Loaded:     Ready
--------------------------------------------------
`
	fmt.Printf(customBanner, appName, address)
}

func GetServerAddress() string {
	host := GetEnv("APP_HOST", "")
	port := GetEnv("APP_PORT", "8080")
	return fmt.Sprintf("%s:%s", host, port)
}
