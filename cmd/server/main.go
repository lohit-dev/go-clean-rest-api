package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lohit-dev/go-clean-rest-api/config"
	"github.com/lohit-dev/go-clean-rest-api/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env", err)
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	app := server.New(cfg.Port, cfg.Logger)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
