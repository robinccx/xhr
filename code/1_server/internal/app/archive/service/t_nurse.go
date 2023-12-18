// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-12-01 20:01:44
// 生成路径: internal/app/archive/service/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================

package service

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"
)

type INurse interface {
	List(ctx context.Context, req *archive.NurseSearchReq) (res *archive.NurseSearchRes, err error)
	GetByNurseId(ctx context.Context, NurseId int64) (res *model.NurseInfoRes, err error)
	Add(ctx context.Context, req *archive.NurseAddReq) (err error)
	Edit(ctx context.Context, req *archive.NurseEditReq) (err error)
	Delete(ctx context.Context, NurseId []int64) (err error)
	Import(ctx context.Context, req *archive.NurseImportReq) (res *archive.NurseImportRes, err error)
	UpdateByImport(ctx context.Context, req *archive.NurseUpdateByImportReq) (res *archive.NurseImportRes, err error)
}

var localNurse INurse

func Nurse() INurse {
	if localNurse == nil {
		panic("implement not found for interface INurse, forgot register?")
	}
	return localNurse
}

func RegisterNurse(i INurse) {
	localNurse = i
}
