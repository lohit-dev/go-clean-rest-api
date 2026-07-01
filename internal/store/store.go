package store

import (
	"errors"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func New(dsn string) (*Store, error) {
	if strings.TrimSpace(dsn) == "" {
		return nil, errors.New("database dsn is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(30 * time.Minute)
	sqlDb.SetConnMaxIdleTime(5 * time.Minute)

	if err := sqlDb.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) DB() *gorm.DB {
	return s.db
}

func (s *Store) Migrate(models ...interface{}) error {
	return s.db.AutoMigrate(models...)
}

func (s *Store) Close() error {
	sqlDb, err := s.db.DB()
	if err != nil {
		return err
	}

	return sqlDb.Close()
}
