package models

import "gorm.io/gorm"

type UserDao struct{ baseDAO }

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"` // Team Member, HR/Admin
}
