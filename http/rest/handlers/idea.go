package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/pkg/models"
)

// CreateIdea handles POST /ideas to create a new idea
type IdeaSubmission struct {
	Input models.Idea
	C     *gin.Context
}

// CreateIdea handles the idea submission
func (r Router) CreateIdea(c *gin.Context) {
	var input models.Idea
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new Idea
	idea := models.Idea{
		Title:            input.Title,
		Description:      input.Description,
		Tags:             input.Tags,
		Timeline:         input.Timeline,
		ImpactEstimation: input.ImpactEstimation,
		UserID:           input.UserID,
	}

	// Save the idea to the database
	if err := r.dao.Idea().Create(&idea); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": idea})
}

// GetIdeas handles GET /ideas to fetch all ideas with pagination
func (r Router) GetIdeas(c *gin.Context) {
	// Get pagination parameters
	page := c.Query("page")
	limit := c.Query("limit")

	// Set default values for pagination
	const defaultLimit = 10
	if limit == "" {
		limit = "10" // Default limit
	}

	// Use the provided page and limit, convert to integers
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1 // Default to page 1 if invalid
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 {
		limitNum = defaultLimit // Default to 10 if invalid
	}

	// Calculate offset for pagination
	offset := (pageNum - 1) * limitNum

	// Fetch total count of ideas
	total, err := r.dao.Idea().GetCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch paginated ideas
	ideas, err := r.dao.Idea().GetPaginated(offset, limitNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Initialize a slice to hold ideas with user names
	ideasWithUserNames := []gin.H{}

	// Fetch the user_name for each idea based on UserID
	for _, idea := range ideas {
		// Fetch user details associated with the idea
		user, err := r.dao.User().ByID(idea.UserID)
		if err != nil {
			// If user is not found, handle error appropriately
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found for idea"})
			return
		}

		// Append the idea and the user_name to the new list
		ideasWithUserNames = append(ideasWithUserNames, gin.H{
			"id":                 idea.ID,
			"title":              idea.Title,
			"description":        idea.Description,
			"tags":               idea.Tags,
			"timeline":           idea.Timeline,          // Custom date format
			"impact_estimation":  idea.ImpactEstimation,
			"user_id":            idea.UserID,
			"user_name":          user.Username,  // Add user_name to the response
			"status":             idea.Status,
			"created_at":         idea.CreatedAt,
			"updated_at":         idea.UpdatedAt,
			"vote_count":         idea.VoteCount,  // Include vote count
		})
	}

	// Create response structure with pagination info
	response := gin.H{
		"total": total,
		"page":  pageNum,
		"limit": limitNum,
		"ideas": ideasWithUserNames,
	}

	// Return the JSON response
	c.JSON(http.StatusOK, response)
}

// GetIdea handles GET /ideas/:id to fetch a specific idea by ID
func (r Router) GetIdea(c *gin.Context) {
	id := c.Param("id")
	idea, err := r.dao.Idea().ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idea not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": idea})
}
