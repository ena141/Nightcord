package main

import (
	"Nightcord/server/internal/api"
	"Nightcord/server/internal/config"
	"Nightcord/server/internal/repository"
	"Nightcord/server/internal/services"
)

func initDI(ds *dataSource, cfg *config.Config) *api.Services {
	userRepository := repository.NewUserRepository(ds.DB)

	userService := services.NewUserService(userRepository)

	return &api.Services{UserService: userService}
}
