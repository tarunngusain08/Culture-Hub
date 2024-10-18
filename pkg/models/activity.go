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

func (a ActivityDao) GetActivities() ([]Activity, error) {
	var activities []Activity
	if err := a.db.Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}
