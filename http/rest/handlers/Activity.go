package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r Router) GetActivities(c *gin.Context) {
	activities, err := r.dao.Activity().GetActivities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"activities": activities})
}
