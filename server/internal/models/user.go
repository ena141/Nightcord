package models

import (
	"gorm.io/gorm"
)

// User models
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex" json:"username"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"password"`
}
