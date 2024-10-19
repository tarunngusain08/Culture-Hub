package models

import "gorm.io/gorm"

type UserDao struct{ baseDAO }

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`        // Team Member, HR/Admin
	Email    string `gorm:"unique;not null"` // Team Member, HR/Admin
}

func (u UserDao) GetByEmail(email string) (*User, error) {
	var user User
	err := u.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u UserDao) GetByUsername(username string) (*User, error) {
	var user User
	err := u.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (u UserDao) ByID(id int) (*User, error) {
	var user User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u UserDao) Create(user *User) error {
	return u.db.Create(user).Error
}
