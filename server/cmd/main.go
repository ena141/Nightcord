package main

import (
	"Nightcord/server/internal/api"
	"Nightcord/server/internal/config"
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// Load env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Set the config
	ctx := context.Background()
	cfg, err := config.LoadConfig(ctx)

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Init Data Source
	ds, err := initDS(cfg)

	if err != nil {
		log.Fatalf("Error initializing datastore: %v", err)
	}

	// Set the dependency & data source injection
	services := initDI(ds, cfg)

	// Get the router
	router := api.NewRouter(services, cfg)

	// Start the server
	go func() {
		if err = http.ListenAndServe(":8080", router); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

}
