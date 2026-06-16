package config

import (
	"os"

	"go.uber.org/zap"
)

type Config struct {
	Port   string
	Logger *zap.Logger
}

func New() (*Config, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:   getEnv("PORT", "4545"),
		Logger: logger,
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
