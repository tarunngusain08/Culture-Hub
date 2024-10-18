package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarunngusain08/Culture-Hub/controller"
	"github.com/tarunngusain08/Culture-Hub/database"
	"github.com/tarunngusain08/Culture-Hub/middleware"
	"github.com/tarunngusain08/Culture-Hub/repo"
	"github.com/tarunngusain08/Culture-Hub/service"
)

type Controller struct {
	*controller.RegisterController
	*controller.LoginController
	*controller.ActivityController
}

var handler *Controller

func main() {
	r := gin.Default()

	// Routes
	r.POST("/api/v1/register", handler.Register)
	r.POST("/api/v1/login", handler.Login)

	// Protected Routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/activities", handler.GetActivities)

	r.Run(":8080")
}

func init() {
	handler = new(Controller)

	// Initialize Database
	db := database.Connect()

	registerRepo := repo.NewRegisterRepo(db)
	loginRepo := repo.NewLoginRepo(db)
	activitiesRepo := repo.NewActivitiesRepo(db)

	registerService := service.NewRegisterService(registerRepo)
	loginService := service.NewLoginService(loginRepo)
	activityService := service.NewActivityService(activitiesRepo)

	registerHandler := controller.NewRegisterController(registerService)
	loginHandler := controller.NewLoginController(loginService)
	activityHandler := controller.NewActivityController(activityService)

	handler.RegisterController = registerHandler
	handler.LoginController = loginHandler
	handler.ActivityController = activityHandler
}
