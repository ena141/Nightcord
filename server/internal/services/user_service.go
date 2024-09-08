package services

import (
	"Nightcord/server/internal/repository"
)

type UserService interface {
	GetUser(uid string)
}

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (*UserServiceImpl) GetUser(uid string) {

}
