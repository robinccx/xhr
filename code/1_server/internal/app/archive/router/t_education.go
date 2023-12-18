// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2023-11-30 18:32:48
// 生成路径: internal/app/archive/router/t_education.go
// 生成人：xiao
// desc:学历
// company:云南奇讯科技有限公司
// ==========================================================================


package router


import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/archive/controller"
)
func (router *Router) BindEducationController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/education", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Education,
		)
	})
}
