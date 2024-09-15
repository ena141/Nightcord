package repository

import (
	"Nightcord/server/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(*models.User) (*models.User, error)
	// UpdateUser(*models.User) (*models.User, error)
	FindByEmail(string) (*models.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) SaveUser(user *models.User) (*models.User, error) {
	err := r.DB.Create(user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
