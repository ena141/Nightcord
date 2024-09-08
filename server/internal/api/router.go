package api

import (
	"Nightcord/server/internal/config"
	"Nightcord/server/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter(services *services.Services, cfg *config.Config) *gin.Engine {
	services.UserService.GetUser(cfg.DatabaseUrl)

	// init router
	router := gin.Default()

	return router
}
