// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/controller/t_department.go
// 生成人：xiao
// desc:科室
// company:云南奇讯科技有限公司
// ==========================================================================


package controller


import (
    "context"
    "github.com/tiger1103/gfast/v3/api/v1/archive"
    "github.com/tiger1103/gfast/v3/internal/app/archive/service"    
    systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"    
)


type departmentController struct {    
    systemController.BaseController    
}


var Department = new(departmentController)


// List 列表
func (c *departmentController) List(ctx context.Context, req *archive.DepartmentSearchReq) (res *archive.DepartmentSearchRes, err error) {
	res, err = service.Department().List(ctx, req)
	return
}


// Get 获取科室
func (c *departmentController) Get(ctx context.Context, req *archive.DepartmentGetReq) (res *archive.DepartmentGetRes, err error) {
    res = new(archive.DepartmentGetRes)
	res.DepartmentInfoRes,err = service.Department().GetByDeptId(ctx, req.DeptId)
	return
}


// Add 添加科室
func (c *departmentController) Add(ctx context.Context, req *archive.DepartmentAddReq) (res *archive.DepartmentAddRes, err error) {
	err = service.Department().Add(ctx, req)
	return
}


// Edit 修改科室
func (c *departmentController) Edit(ctx context.Context, req *archive.DepartmentEditReq) (res *archive.DepartmentEditRes, err error) {
	err = service.Department().Edit(ctx, req)
	return
}


// Delete 删除科室
func (c *departmentController) Delete(ctx context.Context, req *archive.DepartmentDeleteReq) (res *archive.DepartmentDeleteRes, err error) {
	err = service.Department().Delete(ctx, req.DeptIds)
	return
}