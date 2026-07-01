package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lohit-dev/go-clean-rest-api/config"
	"github.com/lohit-dev/go-clean-rest-api/internal/server"
	"github.com/lohit-dev/go-clean-rest-api/internal/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	cfg, err := config.New()
	if err != nil {
		return err
	}
	defer cfg.Logger.Sync()

	db, err := store.New(cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	app := server.New(cfg.Port, cfg.Logger, db)
	if err := app.Run(); err != nil {
		return err
	}

	return nil
}
