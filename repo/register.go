package repo

import (
	"github.com/jinzhu/gorm"
	"github.com/tarunngusain08/Culture-Hub/models"
)

type RegisterRepo struct {
	db *gorm.DB
}

func NewRegisterRepo(db *gorm.DB) *RegisterRepo {
	return &RegisterRepo{db: db}
}

func (r *RegisterRepo) Register(user *models.User, hashedPassword string) error {
	user.Password = string(hashedPassword)
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
