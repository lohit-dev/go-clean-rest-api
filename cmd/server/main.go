package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/lohit-dev/go-clean-rest-api/config"
	"github.com/lohit-dev/go-clean-rest-api/internal/server"
	"github.com/lohit-dev/go-clean-rest-api/internal/store"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env", err)
	}

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("db connect failed: ", err)
	}
	defer db.Close()

	app := server.New(cfg.Port, cfg.Logger, db)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
