package bootstrap

import (
	"go-framework/app/models"
	"go-framework/config"
	"log"
	"log/slog"

	"github.com/golobby/container/v3"
	"gorm.io/gorm"
)

func Migrate() {
	config.RegistererDatabase()
	var db *gorm.DB

	err := container.Resolve(&db)
	if err != nil {
		slog.Error("error when resolve db", "error", err)
	}

	log.Printf("Running auto migrate on all models...")

	for _, model := range models.ModelRegistry {
		err = db.AutoMigrate(model)
		if err != nil {
			slog.Error("error when migrate model", "error", err)
			log.Fatalf("failed to migrate model %v", err)
		}

		log.Printf("model %v", model)
	}

	log.Printf("auto migrate completed successfully")
}
