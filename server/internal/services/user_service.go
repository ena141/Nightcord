package services

import (
	"Nightcord/server/internal/models"
	"Nightcord/server/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *models.User) (*models.User, error)
	VerifyUser(email, password string) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (s *userService) RegisterUser(user *models.User) (*models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return s.userRepository.SaveUser(user)
}

func (s *userService) VerifyUser(email, password string) (bool, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil, nil
}
