package models

import (
	"time"

	"gorm.io/gorm"
)

type ActivityDao struct {
	baseDAO
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
