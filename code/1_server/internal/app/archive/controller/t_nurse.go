// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-12-01 20:01:44
// 生成路径: internal/app/archive/controller/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================

package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type nurseController struct {
	systemController.BaseController
}

var Nurse = new(nurseController)

// List 列表
func (c *nurseController) List(ctx context.Context, req *archive.NurseSearchReq) (res *archive.NurseSearchRes, err error) {
	res, err = service.Nurse().List(ctx, req)
	return
}

// Get 获取护士
func (c *nurseController) Get(ctx context.Context, req *archive.NurseGetReq) (res *archive.NurseGetRes, err error) {
	res = new(archive.NurseGetRes)
	res.NurseInfoRes, err = service.Nurse().GetByNurseId(ctx, req.NurseId)
	return
}

// Add 添加护士
func (c *nurseController) Add(ctx context.Context, req *archive.NurseAddReq) (res *archive.NurseAddRes, err error) {
	err = service.Nurse().Add(ctx, req)
	return
}

// Edit 修改护士
func (c *nurseController) Edit(ctx context.Context, req *archive.NurseEditReq) (res *archive.NurseEditRes, err error) {
	err = service.Nurse().Edit(ctx, req)
	return
}

// Delete 删除护士
func (c *nurseController) Delete(ctx context.Context, req *archive.NurseDeleteReq) (res *archive.NurseDeleteRes, err error) {
	err = service.Nurse().Delete(ctx, req.NurseIds)
	return
}

// 导入
func (c *nurseController) Import(ctx context.Context, req *archive.NurseImportReq) (res *archive.NurseImportRes, err error) {
	res, err = service.Nurse().Import(ctx, req)
	return
}

// 导入
func (c *nurseController) UpdateByImport(ctx context.Context, req *archive.NurseUpdateByImportReq) (res *archive.NurseImportRes, err error) {
	res, err = service.Nurse().UpdateByImport(ctx, req)
	return
}
