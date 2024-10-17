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

	r.Run(":8080")
}
