/*
* @desc:路由绑定
* @company:云南奇讯科技有限公司
* @Author: yixiaohu
* @Date:   2022/2/18 16:23
 */

package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	archiveRouter "github.com/tiger1103/gfast/v3/internal/app/archive/router"
	commonRouter "github.com/tiger1103/gfast/v3/internal/app/common/router"
	reportRouter "github.com/tiger1103/gfast/v3/internal/app/report/router"
	systemRouter "github.com/tiger1103/gfast/v3/internal/app/system/router"
	"github.com/tiger1103/gfast/v3/library/libRouter"
)

var R = new(Router)

type Router struct{}

func (router *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//跨域处理，安全起见正式环境请注释该行
		//group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 绑定后台路由
		systemRouter.R.BindController(ctx, group)
		// 绑定公共路由
		commonRouter.R.BindController(ctx, group)

		// 绑定archive路由
		archiveRouter.R.BindController(ctx, group)

		// 绑定report路由
		reportRouter.R.BindController(ctx, group)

		//自动绑定定义的模块
		if err := libRouter.RouterAutoBind(ctx, router, group); err != nil {
			panic(err)
		}
	})
}
