package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/database"
	"github.com/tarunngusain08/Culture-Hub/models"
)

// CreateIdea handles POST /ideas to create a new idea
func CreateIdea(c *gin.Context) {
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
	if err := database.DB.Create(&idea).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": idea})
}

// GetIdeas handles GET /ideas to fetch all ideas with pagination
func GetIdeas(c *gin.Context) {
	var ideas []models.Idea
	var total int64

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
	database.DB.Model(&models.Idea{}).Count(&total)

	// Fetch paginated ideas
	if err := database.DB.Offset(offset).Limit(limitNum).Find(&ideas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create response structure with pagination info
	response := gin.H{
		"total": total,
		"page":  pageNum,
		"limit": limitNum,
		"ideas": ideas,
	}

	c.JSON(http.StatusOK, response)
}

// GetIdea handles GET /ideas/:id to fetch a specific idea by ID
func GetIdea(c *gin.Context) {
	var idea models.Idea
	id := c.Param("id")

	if err := database.DB.First(&idea, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idea not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": idea})
}
