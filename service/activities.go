package service

import (
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/repo"
)

type ActivityService struct {
	activityRepo *repo.ActivityRepo
}

func NewActivityService(activityRepo *repo.ActivityRepo) *ActivityService {
	return &ActivityService{activityRepo}
}

func (a *ActivityService) GetActivities() ([]models.Activity, error) {
	return a.activityRepo.GetActivities()
}
