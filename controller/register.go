package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/models"
	"github.com/tarunngusain08/Culture-Hub/service"
	"net/http"
)

type RegisterController struct {
	registerService *service.RegisterService
}

func NewRegisterController(registerService *service.RegisterService) *RegisterController {
	return &RegisterController{registerService: registerService}
}

func (r *RegisterController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := r.registerService.Register(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}
