package rest

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/config"
	"github.com/tarunngusain08/culturehub/http/rest/handlers"
	"github.com/tarunngusain08/culturehub/http/rest/middleware"
	"github.com/tarunngusain08/culturehub/http/rest/session"
	"github.com/tarunngusain08/culturehub/pkg/models"
)

type Server struct {
	server *gin.Engine
}

func newServer() *Server {
	return &Server{server: gin.Default()}
}

// // Serve starts the server
func Serve(dao models.DaoService) error {
	s := newServer()

	router := handlers.NewRouter(context.Background(), dao)
	s.middlewares(dao)
	s.routing(router)

	return s.server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (s *Server) middlewares(dao models.DaoService) {
	s.server.Use(middleware.CorsMiddleware)
	// s.server.Use(middleware.RateLimiter)
	s.server.Use(middleware.Logger)
	s.server.Use(session.Middleware)
	s.server.Use(middleware.BasicAuthMiddleware(dao))
	// s.server.Use(middleware.PanicRecover)
}

func (s *Server) routing(r *handlers.Router) {
	s.server.Static("/assets", config.GetAppPath()+"http/assets/")

	s.server.GET("/hello/world", r.HelloWorldHandler)

	s.server.POST(handlers.LoginPath)
	s.server.POST("/api/v1/register")
	s.server.POST("/api/v1/ideas")
	s.server.GET("/api/v1/ideas")
	s.server.GET("/api/v1/ideas/{id}")

	s.server.POST("/api/v1/ideas/:id/vote", r.VoteCountHandler)

	s.server.NoRoute()
}
