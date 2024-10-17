package handlers

import (
	"context"

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
