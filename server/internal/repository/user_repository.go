package repository

import "gorm.io/gorm"

type UserRepository interface {
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: DB,
	}
}

func GetUser(uid string) {

}
