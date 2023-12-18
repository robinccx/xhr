// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/controller/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================


package controller


import (
    "context"
    "github.com/tiger1103/gfast/v3/api/v1/archive"
    "github.com/tiger1103/gfast/v3/internal/app/archive/service"    
    systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"    
)


type nurseChangeLogController struct {    
    systemController.BaseController    
}


var NurseChangeLog = new(nurseChangeLogController)


// List 列表
func (c *nurseChangeLogController) List(ctx context.Context, req *archive.NurseChangeLogSearchReq) (res *archive.NurseChangeLogSearchRes, err error) {
	res, err = service.NurseChangeLog().List(ctx, req)
	return
}


// Get 获取变更记录查询
func (c *nurseChangeLogController) Get(ctx context.Context, req *archive.NurseChangeLogGetReq) (res *archive.NurseChangeLogGetRes, err error) {
    res = new(archive.NurseChangeLogGetRes)
	res.NurseChangeLogInfoRes,err = service.NurseChangeLog().GetById(ctx, req.Id)
	return
}


// Add 添加变更记录查询
func (c *nurseChangeLogController) Add(ctx context.Context, req *archive.NurseChangeLogAddReq) (res *archive.NurseChangeLogAddRes, err error) {
	err = service.NurseChangeLog().Add(ctx, req)
	return
}


// Edit 修改变更记录查询
func (c *nurseChangeLogController) Edit(ctx context.Context, req *archive.NurseChangeLogEditReq) (res *archive.NurseChangeLogEditRes, err error) {
	err = service.NurseChangeLog().Edit(ctx, req)
	return
}


// Delete 删除变更记录查询
func (c *nurseChangeLogController) Delete(ctx context.Context, req *archive.NurseChangeLogDeleteReq) (res *archive.NurseChangeLogDeleteRes, err error) {
	err = service.NurseChangeLog().Delete(ctx, req.Ids)
	return
}