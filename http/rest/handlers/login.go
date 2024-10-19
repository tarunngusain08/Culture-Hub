package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/tarunngusain08/culturehub/http/rest/utils"
	"github.com/tarunngusain08/culturehub/pkg/models"
)

const LoginPath = "/api/v1/login"

func (r Router) Login(c *gin.Context) {
	var userDetails models.User
	if err := c.ShouldBindJSON(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.isValidUser(&userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := token(&userDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func token(userDetails *models.User) (string, error) {
	token, err := utils.GenerateToken(userDetails.ID, userDetails.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *Router) isValidUser(userDetails *models.User) error {
	var user *models.User
	var err error

	isEmail := utils.IsValidEmail(userDetails.Email)

	if isEmail {
		user, err = r.dao.User().GetByEmail(userDetails.Email)
		if err != nil {
			return err
		}
	} else {
		user, err = r.dao.User().GetByUsername(userDetails.Username)
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDetails.Password)); err != nil {
		return err
	}

	return nil
}
