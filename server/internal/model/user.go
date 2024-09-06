package model

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
}
