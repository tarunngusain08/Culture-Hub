package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/service"
	"net/http"
)

type ActivityController struct {
	activityService *service.ActivityService
}

func NewActivityController(activityService *service.ActivityService) *ActivityController {
	return &ActivityController{activityService: activityService}
}

func (a *ActivityController) GetActivities(c *gin.Context) {
	activities, err := a.activityService.GetActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"activities": activities})
}
