package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/http/components"
	"github.com/tarunngusain08/culturehub/http/rest/gintemplrenderer"
)

func (r *Router) HelloWorldHandler(c *gin.Context) {
	r.RenderHelloworld(c, 200, components.FormPage, nil)
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func (r Router) RenderHelloworld(ctx *gin.Context, statusCode int, t func(templ.Component) templ.Component, ideas []components.Idea) {
	fmt.Println("ideas!!", ideas)
	rc := gintemplrenderer.New(ctx.Request.Context(), http.StatusOK, t(components.IdeaList(ideas)))
	ctx.Render(http.StatusOK, rc)
}
