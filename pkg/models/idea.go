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

func (IdeaDao) TableName() string {
	return "ideas"
}

func (i IdeaDao) ByID(id string) (*Idea, error) {
	var idea Idea
	err := i.db.Where("id =?", id).First(&idea).Error
	return &idea, err
}

func (i IdeaDao) VoteUpdate(id string, vote int) error {
	var err error
	if vote > 0 {
		err = i.db.Model(&Idea{}).Where("id = ?", id).Update("vote_count", gorm.Expr("vote_count - ?", 1)).Error
	} else {
		err = i.db.Model(&Idea{}).Where("id = ?", id).Update("vote_count", gorm.Expr("vote_count - ?", 1)).Error
	}
	return err
}

func (i IdeaDao) Create(idea *Idea) error {
	return i.baseDAO.db.Create(idea).Error
}

func (i IdeaDao) GetCount() (count int64, err error) {
	err = i.baseDAO.db.Model(&Idea{}).Count(&count).Error
	return
}

func (i IdeaDao) GetPaginated(offset, limit int) ([]Idea, error) {
	var ideas []Idea
	err := i.db.Offset(offset).Limit(limit).Find(&ideas).Error
	return ideas, err
}
