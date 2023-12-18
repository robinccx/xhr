package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/archive/controller"
)

func (router *Router) BindImportController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/import", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Title,
		)
	})
}