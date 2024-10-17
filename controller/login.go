package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/service"
	"net/http"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(loginService *service.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (l *LoginController) Login(c *gin.Context) {
	var userDetails models.User
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := l.loginService.Login(&userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
