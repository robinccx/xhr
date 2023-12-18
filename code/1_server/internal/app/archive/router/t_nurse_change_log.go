// ==========================================================================
// GFast自动生成router操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/router/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================


package router


import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast/v3/internal/app/archive/controller"
)
func (router *Router) BindNurseChangeLogController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/nurseChangeLog", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.NurseChangeLog,
		)
	})
}
