package config

import "go-framework/app/http/services"

func RegisterService() {
	services.AutoRegisterService()
}

func init() {
	RegisterConfig(func() error {
		RegisterService()
		return nil
	})
}
