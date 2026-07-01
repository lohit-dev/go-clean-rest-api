package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lohit-dev/go-clean-rest-api/config"
	"github.com/lohit-dev/go-clean-rest-api/internal/auth"
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

	if err := db.Migrate(
		&auth.User{},
		&auth.Session{},
		&auth.OTPCode{},
		&auth.Passkey{},
		&auth.OAuthAccount{},
	); err != nil {
		return err
	}

	log.Println("migrations complete")

	return nil
}
