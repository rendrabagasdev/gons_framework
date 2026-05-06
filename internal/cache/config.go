package cache

import (
	"gons/internal/contracts"
	"gons/internal/registry"

	"github.com/golobby/container/v3"
)

func RegisterCache() {
	container.Singleton(func() contracts.Cache {
		return NewCache()
	})
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterCache()
		return nil
	})
}
