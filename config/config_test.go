package config

import "testing"

func TestNewReturnsErrorWhenDatabaseURLMissing(t *testing.T) {
	t.Setenv("PORT", "8080")
	t.Setenv("DB_URL", "")

	cfg, err := New()
	if err == nil {
		t.Fatalf("expected error, got nil and config %#v", cfg)
	}
}

func TestNewUsesDefaultsAndEnvironmentValues(t *testing.T) {
	t.Setenv("DB_URL", "postgres://user:pass@localhost:5432/app?sslmode=disable")
	t.Setenv("PORT", "")

	cfg, err := New()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if cfg.Port != "4545" {
		t.Fatalf("expected default port 4545, got %s", cfg.Port)
	}

	if cfg.DatabaseURL == "" {
		t.Fatal("expected database url to be set")
	}
}
