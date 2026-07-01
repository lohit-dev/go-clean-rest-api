package config

import (
	"errors"
	"os"
	"strings"

	"go.uber.org/zap"
)

type Config struct {
	Port        string
	Logger      *zap.Logger
	DatabaseURL string
}

func New() (*Config, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:        getEnv("PORT", "4545"),
		Logger:      logger,
		DatabaseURL: getEnv("DB_URL", ""),
	}

	if err := cfg.validate(); err != nil {
		_ = logger.Sync()
		return nil, err
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func (c *Config) validate() error {
	if strings.TrimSpace(c.DatabaseURL) == "" {
		return errors.New("missing required environment variable: DB_URL")
	}

	if strings.TrimSpace(c.Port) == "" {
		return errors.New("missing server port")
	}

	return nil
}
