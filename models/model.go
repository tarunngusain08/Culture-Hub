package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"` // Team Member, HR/Admin
}

type Activity struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      string `gorm:"not null"` // Ongoing, Upcoming, Completed
	StartDate   time.Time
	EndDate     time.Time
	Votes       int
	CreatedBy   uint
}
