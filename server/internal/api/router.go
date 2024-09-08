package api

import (
	"Nightcord/server/internal/config"
	"Nightcord/server/internal/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userHandler *UserHandler
}

func SetupRouter(services *services.Services, cfg *config.Config) *gin.Engine {
	userHandler := NewUserHandler(services)
	handler := &Handler{
		userHandler: userHandler,
	}

	// init router
	router := gin.Default()

	router.Group("/user")
	{
		router.GET("/:id", handler.userHandler.GetUser)
	}

	return router
}
