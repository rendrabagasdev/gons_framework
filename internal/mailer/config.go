package mailer

import (
	"gons/internal/contracts"
	"gons/internal/registry"
	"log/slog"

	"github.com/golobby/container/v3"
)

func RegisterMailer() {
	err := container.Singleton(func() contracts.Mailer {
		return NewMailer()
	})

	if err != nil {
		slog.Error("Gons: Mailer register error: " + err.Error())
	}
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterMailer()
		return nil
	})
}
