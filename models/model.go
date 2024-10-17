package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"` // Team Member, HR/Admin
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

type Idea struct {
	gorm.Model
	Title            string    `gorm:"type:varchar(255);not null" json:"title"`
	Description      string    `gorm:"type:text;not null" json:"description"`
	Tags             []string  `gorm:"type:text[]" json:"tags"` // Array of strings
	Timeline         time.Time `json:"timeline"`
	ImpactEstimation string    `json:"impact_estimation"`
	UserID           uint      `json:"user_id"`                                            // Foreign key to User table
	Status           string    `gorm:"type:idea_status;default:'Submitted'" json:"status"` // Default to 'Submitted'
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	VoteCount        int       `gorm:"default:0" json:"vote_count"` // Default vote count
}
