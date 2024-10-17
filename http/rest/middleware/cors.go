package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	config = cors.DefaultConfig()
	config.AllowAllOrigins = true
	CorsMiddleware = cors.New(config)
}

var (
	config         cors.Config
	CorsMiddleware gin.HandlerFunc
)
