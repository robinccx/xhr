// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-11-30 18:32:48
// 生成路径: internal/app/archive/service/t_education.go
// 生成人：xiao
// desc:学历
// company:云南奇讯科技有限公司
// ==========================================================================


package service


import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"	
)


type IEducation interface {
	List(ctx context.Context, req *archive.EducationSearchReq) (res *archive.EducationSearchRes, err error)
	GetByEducationId(ctx context.Context, EducationId int64) (res *model.EducationInfoRes,err error)
	Add(ctx context.Context, req *archive.EducationAddReq) (err error)
	Edit(ctx context.Context, req *archive.EducationEditReq) (err error)
	Delete(ctx context.Context, EducationId []int64) (err error)
}


var localEducation IEducation


func Education() IEducation {
	if localEducation == nil {
		panic("implement not found for interface IEducation, forgot register?")
	}
	return localEducation
}


func RegisterEducation(i IEducation) {
	localEducation = i
}
