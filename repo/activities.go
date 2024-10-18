package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tarunngusain08/Culture-Hub/models"
)

type ActivityRepo struct {
	db *gorm.DB
}

func NewActivitiesRepo(db *gorm.DB) *ActivityRepo {
	return &ActivityRepo{db: db}
}

func (a *ActivityRepo) GetActivities() ([]models.Activity, error) {
	var activities []models.Activity
	if err := a.db.Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}
