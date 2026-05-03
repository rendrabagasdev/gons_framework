package seeders

import (
	"log/slog"

	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

type Seeder func(db *gorm.DB) error

var SeederRegistry []Seeder

func RegisterSeeder(seeder Seeder) {
	SeederRegistry = append(SeederRegistry, seeder)
}

func RunSeeders(db *gorm.DB) {
	log.Info("Starting database seeding...")

	for _, seeder := range SeederRegistry {
		if err := seeder(db); err != nil {
			slog.Error("Seeder failed to execute", "error", err)
			return
		}
	}

	log.Info("Database Seeding Completed!")
}
