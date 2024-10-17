package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tarunngusain08/Culture-Hub/models"
	"golang.org/x/crypto/bcrypt"
)

type ActivitiesRepo struct {
	db *gorm.DB
}

func NewActivitiesRepo(db *gorm.DB) *ActivitiesRepo {
	return &ActivitiesRepo{db: db}
}

func (l *ActivitiesRepo) GetActivities(userDetails *models.User) error {
	var user models.User
	if err := l.db.Where("username = ?", userDetails.Username).First(&user).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password)); err != nil {
		return err
	}

	return nil
}
