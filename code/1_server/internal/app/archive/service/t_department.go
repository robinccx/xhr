// ==========================================================================
// GFast自动生成service操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/service/t_department.go
// 生成人：xiao
// desc:科室
// company:云南奇讯科技有限公司
// ==========================================================================


package service


import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"	
)


type IDepartment interface {
	List(ctx context.Context, req *archive.DepartmentSearchReq) (res *archive.DepartmentSearchRes, err error)
	GetByDeptId(ctx context.Context, DeptId int64) (res *model.DepartmentInfoRes,err error)
	Add(ctx context.Context, req *archive.DepartmentAddReq) (err error)
	Edit(ctx context.Context, req *archive.DepartmentEditReq) (err error)
	Delete(ctx context.Context, DeptId []int64) (err error)
}


var localDepartment IDepartment


func Department() IDepartment {
	if localDepartment == nil {
		panic("implement not found for interface IDepartment, forgot register?")
	}
	return localDepartment
}


func RegisterDepartment(i IDepartment) {
	localDepartment = i
}
