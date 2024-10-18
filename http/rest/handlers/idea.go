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

// Create a channel for the queue
var ideaQueue = make(chan IdeaSubmission, 100) // Adjust the buffer size as needed, Process 100 message to queue simultaneously

// Worker function to process idea submissions
func ideaWorker(dao models.DaoService) {
	for submission := range ideaQueue {
		input := submission.Input
		c := submission.C

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
		if err := dao.Idea().Create(&idea); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			continue // Skip to the next submission on error
		}

		c.JSON(http.StatusOK, gin.H{"data": idea})
	}
}

// CreateIdea handles the idea submission
func (r Router) CreateIdea(c *gin.Context) {
	var input models.Idea
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send the idea submission to the queue
	ideaQueue <- IdeaSubmission{Input: input, C: c}
}

// Initialize function to start the worker
func Init(dao models.DaoService) {
	go ideaWorker(dao) // Start the worker in a goroutine
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
func (r Router) GetIdea(c *gin.Context) {
	id := c.Param("id")
	idea, err := r.dao.Idea().ByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Idea not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": idea})
}
