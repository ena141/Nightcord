package api

import (
	"Nightcord/server/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(services *services.Services) *UserHandler {
	return &UserHandler{userService: services.UserService}
}

func (UserHandler *UserHandler) GetUser(c *gin.Context) {
	UserHandler.userService.GetUser("123")
}
