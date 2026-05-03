package bootstrap

import (
	"go-framework/config"
	"go-framework/database/seeders"
	"go-framework/routes"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golobby/container/v3"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func RunApp() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}

	config.AutoRegisterConfig()

	if len(os.Args) > 1 && os.Args[1] == "--seed" {
		var db *gorm.DB
		if err := container.Resolve(&db); err != nil {
			log.Fatalf("Fatal: Gagal menarik DB dari container: %v", err)
		}
		seeders.RunSeeders(db)
		os.Exit(0) // Matikan program setelah proses seed selesai!
	}

	var app *fiber.App
	if err := container.Resolve(&app); err != nil {
		slog.Error("Gons: App resolution error", "error", err)
		log.Fatalf("failed to resolve app: %v", err)
	}

	routes.RegisterRoute(app)
	address := config.GetServerAddress()

	config.PrintBanner()
	if err := app.Listen(address); err != nil {
		log.Fatalf("Server crash: %v", err)
	}
}
