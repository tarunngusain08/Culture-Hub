package handlers

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"github.com/tarunngusain08/culturehub/http/rest/gintemplrenderer"
	"github.com/tarunngusain08/culturehub/pkg/models"
)

type Router struct {
	dao models.DaoService
}

func NewRouter(c context.Context, dao models.DaoService) *Router {
	return &Router{
		dao: dao,
	}
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func (r Router) Render(ctx *gin.Context, statusCode int, t func() templ.Component) {
	rc := gintemplrenderer.New(ctx.Request.Context(), http.StatusOK, t())
	ctx.Render(http.StatusOK, rc)
}
