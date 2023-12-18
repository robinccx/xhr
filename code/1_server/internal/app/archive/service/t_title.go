// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-11-30 18:32:51
// 生成路径: internal/app/archive/service/t_title.go
// 生成人：xiao
// desc:职称
// company:云南奇讯科技有限公司
// ==========================================================================


package service


import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"	
)


type ITitle interface {
	List(ctx context.Context, req *archive.TitleSearchReq) (res *archive.TitleSearchRes, err error)
	GetByTitleId(ctx context.Context, TitleId int64) (res *model.TitleInfoRes,err error)
	Add(ctx context.Context, req *archive.TitleAddReq) (err error)
	Edit(ctx context.Context, req *archive.TitleEditReq) (err error)
	Delete(ctx context.Context, TitleId []int64) (err error)
}


var localTitle ITitle


func Title() ITitle {
	if localTitle == nil {
		panic("implement not found for interface ITitle, forgot register?")
	}
	return localTitle
}


func RegisterTitle(i ITitle) {
	localTitle = i
}
