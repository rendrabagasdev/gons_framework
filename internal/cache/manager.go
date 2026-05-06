package cache

import (
	"gons/internal/contracts"
	"gons/pkg/env"
	"gons/pkg/utils/cache"
	"log/slog"

	"github.com/golobby/container/v3"
	"github.com/redis/go-redis/v9"
)

// NewCache returns a new Cache contract implementation based on environment configuration.
func NewCache() contracts.Cache {
	driver := env.Get("CACHE_DRIVER", "memory")

	if driver == "redis" {
		var client *redis.Client
		if err := container.Resolve(&client); err != nil {
			slog.Error("Gons: Failed to resolve Redis client for Cache", "error", err)
			return &cache.MemoryDriver{}
		}
		return &RedisDriver{Client: client}
	}

	return &cache.MemoryDriver{}
}
