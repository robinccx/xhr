package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/report/controller"
)

func (router *Router) BindNurseMonthlyController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/nurse", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.NurseMonthly,
		)
	})
}
