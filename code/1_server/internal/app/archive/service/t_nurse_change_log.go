// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/service/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================


package service


import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"	
)


type INurseChangeLog interface {
	List(ctx context.Context, req *archive.NurseChangeLogSearchReq) (res *archive.NurseChangeLogSearchRes, err error)
	GetById(ctx context.Context, Id int64) (res *model.NurseChangeLogInfoRes,err error)
	Add(ctx context.Context, req *archive.NurseChangeLogAddReq) (err error)
	Edit(ctx context.Context, req *archive.NurseChangeLogEditReq) (err error)
	Delete(ctx context.Context, Id []int64) (err error)
}


var localNurseChangeLog INurseChangeLog


func NurseChangeLog() INurseChangeLog {
	if localNurseChangeLog == nil {
		panic("implement not found for interface INurseChangeLog, forgot register?")
	}
	return localNurseChangeLog
}


func RegisterNurseChangeLog(i INurseChangeLog) {
	localNurseChangeLog = i
}
