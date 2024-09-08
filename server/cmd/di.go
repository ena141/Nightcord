package main

import (
	"Nightcord/server/internal/config"
	"Nightcord/server/internal/repository"
	"Nightcord/server/internal/services"
)

func initDI(ds *dataSource, cfg *config.Config) *services.Services {
	userRepository := repository.NewUserRepository(ds.DB)

	userService := services.NewUserService(userRepository)

	return &services.Services{UserService: userService}
}
