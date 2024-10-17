package handlers

import "github.com/gin-gonic/gin"

func HelloWorld(c *gin.Context) {
	c.Writer.Write([]byte("hello world"))
}
