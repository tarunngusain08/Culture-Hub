package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/database"
	"github.com/tarunngusain08/Culture-Hub/models"
	"net/http"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	if err := database.DB.Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve activities"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})
}
