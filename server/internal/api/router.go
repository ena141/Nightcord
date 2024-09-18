package api

import (
	"Nightcord/server/internal/config"
	"Nightcord/server/internal/services"
	"github.com/gin-gonic/gin"
)

type Services struct {
	UserService services.UserService
	JWTService  services.JWTService
}

func NewRouter(services *Services, cfg *config.Config) *gin.Engine {
	userHandler := NewUserHandler(services)
	authHandler := NewAuthHandler(services)

	// init router
	router := gin.Default()

	userGroup := router.Group("/user")
	{
		userGroup.GET("/:id", userHandler.GetUser)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	return router
}
