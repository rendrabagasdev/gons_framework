package config

import (
	"log"
	"log/slog"
)

type ConfigProvider func() error

var ConfigRegistry []ConfigProvider

func RegisterConfig(provider ConfigProvider) {
	ConfigRegistry = append(ConfigRegistry, provider)
}

func AutoRegisterConfig() {
	for _, provider := range ConfigRegistry {
		if err := provider(); err != nil {
			slog.Error("Failed to register config", "error", err)
			log.Printf("Failed to register config: %v\n", err)
		}
	}
}
