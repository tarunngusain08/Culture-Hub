package service

import (
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/repo"
	"github.com/tarunngusain08/Culture-Hub/utils"
)

type LoginService struct {
	loginRepo *repo.LoginRepo
}

func NewLoginService(loginRepo *repo.LoginRepo) *LoginService {
	return &LoginService{loginRepo}
}

func (l *LoginService) Login(userDetails *models.User) (string, error) {
	err := l.loginRepo.Login(userDetails)
	if err != nil {
		return "", err
	}
	token, err := utils.GenerateToken(userDetails.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
