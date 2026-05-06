package queue

import (
	"gons/internal/contracts"
	"gons/internal/registry"

	"github.com/golobby/container/v3"
)

func RegisterQueue() {
	container.Singleton(func() contracts.Queue {
		return NewQueue()
	})
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterQueue()
		return nil
	})
}
