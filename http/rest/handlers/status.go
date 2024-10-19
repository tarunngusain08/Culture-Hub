package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) UpdateStatus(c *gin.Context) {
	ideaID := c.Param("id")
	action := c.Query("action")

	switch action {
	case "submitted", "ongoing", "approved", "completed":
		r.dao.Idea().UpdateStatus(action, ideaID)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}

	c.JSON(http.StatusOK, ideaID)
}
