package main

import (
	"Nightcord/server/internal/config"
	"Nightcord/server/internal/model"
	"gorm.io/gorm"
	"log"
)

type dataSource struct {
	DB *gorm.DB
}

func initDS(cfg config.Config) (*dataSource, error) {
	log.Println("Initializing data source ...")

	db, err := config.InitDB(cfg)

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&model.User{},
	); err != nil {
		return nil, err
	}

	log.Println("Data source initialized")

	return &dataSource{DB: db}, nil
}
