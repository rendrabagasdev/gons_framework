package registry

import (
	"log/slog"
)

type ConfigProvider func() error

var ConfigRegistry []ConfigProvider

func RegisterConfig(provider ConfigProvider) {
	ConfigRegistry = append(ConfigRegistry, provider)
}

func AutoRegisterConfig() error {
	for _, provider := range ConfigRegistry {
		if err := provider(); err != nil {
			slog.Error("Failed to register config", "error", err)
			return err
		}
	}
	return nil
}
