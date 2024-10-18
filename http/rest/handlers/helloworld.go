package handlers

import "github.com/gin-gonic/gin"

func (r *Router) HelloWorldHandler(c *gin.Context) {
	c.Writer.Write([]byte("hello world"))
}
