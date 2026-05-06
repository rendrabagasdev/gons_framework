package database

import (
	"log/slog"
	"gons/internal/registry"

	"github.com/golobby/container/v3"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterDatabase() {
	err := container.Singleton(func() *gorm.DB {
		return NewGormConnection()
	})

	if err != nil {
		slog.Error("Gons: database register error: " + err.Error())
	}

	err = container.Singleton(func() *redis.Client {
		return NewRedisClient()
	})

	if err != nil {
		slog.Error("Gons: redis register error: " + err.Error())
	}
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterDatabase()
		return nil
	})
}
