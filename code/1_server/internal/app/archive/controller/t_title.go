// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-11-30 18:32:51
// 生成路径: internal/app/archive/controller/t_title.go
// 生成人：xiao
// desc:职称
// company:云南奇讯科技有限公司
// ==========================================================================


package controller


import (
    "context"
    "github.com/tiger1103/gfast/v3/api/v1/archive"
    "github.com/tiger1103/gfast/v3/internal/app/archive/service"    
    systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"    
)


type titleController struct {    
    systemController.BaseController    
}


var Title = new(titleController)


// List 列表
func (c *titleController) List(ctx context.Context, req *archive.TitleSearchReq) (res *archive.TitleSearchRes, err error) {
	res, err = service.Title().List(ctx, req)
	return
}


// Get 获取职称
func (c *titleController) Get(ctx context.Context, req *archive.TitleGetReq) (res *archive.TitleGetRes, err error) {
    res = new(archive.TitleGetRes)
	res.TitleInfoRes,err = service.Title().GetByTitleId(ctx, req.TitleId)
	return
}


// Add 添加职称
func (c *titleController) Add(ctx context.Context, req *archive.TitleAddReq) (res *archive.TitleAddRes, err error) {
	err = service.Title().Add(ctx, req)
	return
}


// Edit 修改职称
func (c *titleController) Edit(ctx context.Context, req *archive.TitleEditReq) (res *archive.TitleEditRes, err error) {
	err = service.Title().Edit(ctx, req)
	return
}


// Delete 删除职称
func (c *titleController) Delete(ctx context.Context, req *archive.TitleDeleteReq) (res *archive.TitleDeleteRes, err error) {
	err = service.Title().Delete(ctx, req.TitleIds)
	return
}