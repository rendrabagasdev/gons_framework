package config

import (
	"go-framework/app/contracts"
	"go-framework/app/utils/cache"

	"github.com/golobby/container/v3"
)

func init() {
	RegisterConfig(func() error {
		return container.Singleton(func() contracts.Cache {
			driver := GetEnv("CACHE_DRIVER", "memory")

			if driver == "redis" {
				// return &cache.RedisDriver{...}
			}

			return &cache.MemoryDriver{}
		})
	})
}
