// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/service/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================


package service


import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"	
)


type IPosition interface {
	List(ctx context.Context, req *archive.PositionSearchReq) (res *archive.PositionSearchRes, err error)
	GetByPositionId(ctx context.Context, PositionId int64) (res *model.PositionInfoRes,err error)
	Add(ctx context.Context, req *archive.PositionAddReq) (err error)
	Edit(ctx context.Context, req *archive.PositionEditReq) (err error)
	Delete(ctx context.Context, PositionId []int64) (err error)
}


var localPosition IPosition


func Position() IPosition {
	if localPosition == nil {
		panic("implement not found for interface IPosition, forgot register?")
	}
	return localPosition
}


func RegisterPosition(i IPosition) {
	localPosition = i
}
