package services

import (
	"Nightcord/server/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserService interface {
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func GetUser(c *gin.Context) {

}
