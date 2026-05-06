package storage

import (
	"gons/internal/contracts"
	"gons/internal/registry"

	"github.com/golobby/container/v3"
)

func RegisterStorage() {
	container.Singleton(func() contracts.Storage {
		return NewStorage()
	})
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterStorage()
		return nil
	})
}
