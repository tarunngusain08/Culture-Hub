package models

import (
	"time"

	"gorm.io/gorm"
)

type IdeaDao struct {
	baseDAO
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
