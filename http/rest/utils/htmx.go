package utils

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/http/rest/gintemplrenderer"
)

type HttpResponse interface{}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx *gin.Context, statusCode int, t func() templ.Component) {
	r := gintemplrenderer.New(ctx.Request.Context(), http.StatusOK, t())
	ctx.Render(http.StatusOK, r)
}
