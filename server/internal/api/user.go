package api

import (
	"Nightcord/server/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(services *Services) *UserHandler {
	return &UserHandler{userService: services.UserService}
}

func (h *UserHandler) GetUser(c *gin.Context) {

}
