package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/controller"
	"github.com/tarunngusain08/Culture-Hub/database"
	"github.com/tarunngusain08/Culture-Hub/middleware"
)

func main() {
	r := gin.Default()

	// Initialize Database
	database.Connect()

	// Routes
	r.POST("/api/v1/register", controller.Register)
	r.POST("/api/v1/login", controller.Login)

	// Protected Routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/activities", controller.GetActivities)
	// Idea Routes (Add these lines for the Idea model)
	controller.Init()
	protected.POST("/ideas", controller.CreateIdea) // Create a new idea
	protected.GET("/ideas", controller.GetIdeas)    // Get all ideas
	protected.GET("/ideas/:id", controller.GetIdea) // Get a specific idea by ID

	r.Run(":8080")
}
