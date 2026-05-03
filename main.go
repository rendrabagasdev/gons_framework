package main

import (
	"go-framework/bootstrap"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Logger Init
	f, _ := os.OpenFile("logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	slog.SetDefault(slog.New(slog.NewTextHandler(f, nil)))
	log.SetOutput(os.Stdout)

	// Load Envss
	if err := godotenv.Load(); err != nil {
		slog.Error("file .env not found", "error", err)
	}

	// Migrate Command
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "migrate":
			log.Println("Running migrate command...")
			bootstrap.Migrate()
			return
		}
	}

	// Bootstrap App
	bootstrap.RunApp()

}
