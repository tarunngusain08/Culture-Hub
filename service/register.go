package service

import (
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/repo"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	registerRepo *repo.RegisterRepo
}

func NewRegisterService(registerRepo *repo.RegisterRepo) *RegisterService {
	return &RegisterService{registerRepo}
}

func (r *RegisterService) Register(user *models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return r.registerRepo.Register(user, hashedPassword)
}
