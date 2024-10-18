package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) VoteCountHandler(c *gin.Context) {
	ideaID := c.Param("id")
	action := c.Query("action")

	switch action {
	case "upvote":
		r.dao.Idea().VoteUpdate(ideaID, 1)
	case "downvote":
		r.dao.Idea().VoteUpdate(ideaID, -1)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}

	c.JSON(http.StatusOK, ideaID)
}
