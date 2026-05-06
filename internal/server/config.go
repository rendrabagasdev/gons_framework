package server

import (
	"fmt"
	"log"
	"log/slog"

	"gons/pkg/env"
	"gons/internal/registry"

	"github.com/gofiber/fiber/v3"
	"github.com/golobby/container/v3"
)

func RegisterServer() {
	err := container.Singleton(func() *fiber.App {
		return NewServer()
	})

	if err != nil {
		slog.Error("error when register server", "error", err)
		log.Fatalf("failed to register server %v", err)
	}
}

func init() {
	registry.RegisterConfig(func() error {
		RegisterServer()
		return nil
	})
}

// PrintBanner mencetak ASCII art Gons ke terminal
func PrintBanner() {
	appName := env.Get("APP_NAME", "Gons Framework")
	address := GetServerAddress()

	customBanner := `
   ______                 
  / ____/___  ____  _____ 
 / / __/ __ \/ __ \/ ___/ 
/ /_/ / /_/ / / / (__  )  
\____/\____/_/ /_/____/   
                          
🚀 %s v1.0.0
--------------------------------------------------
INFO  Server Address:    http://%s
INFO  Environment:       Development
INFO  Routes Loaded:     Ready
--------------------------------------------------
`
	fmt.Printf(customBanner, appName, address)
}

func GetServerAddress() string {
	host := env.Get("APP_HOST", "")
	port := env.Get("APP_PORT", "8080")
	return fmt.Sprintf("%s:%s", host, port)
}
