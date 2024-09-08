package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB(cfg *Config) (*gorm.DB, error) {
	log.Printf("Connecting to Postgres\n")
	db, err := gorm.Open(postgres.Open(cfg.DatabaseUrl))

	if err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to Postgres\n")
	return db, nil
}
