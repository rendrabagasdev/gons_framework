package service

import (
	"gons/app/http/services"
	"gons/internal/registry"
)

func RegisterService() {
	services.AutoRegisterService()
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterService()
		return nil
	})
}
