// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/controller/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================


package controller


import (
    "context"
    "github.com/tiger1103/gfast/v3/api/v1/archive"
    "github.com/tiger1103/gfast/v3/internal/app/archive/service"    
    systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"    
)


type positionController struct {    
    systemController.BaseController    
}


var Position = new(positionController)


// List 列表
func (c *positionController) List(ctx context.Context, req *archive.PositionSearchReq) (res *archive.PositionSearchRes, err error) {
	res, err = service.Position().List(ctx, req)
	return
}


// Get 获取职务
func (c *positionController) Get(ctx context.Context, req *archive.PositionGetReq) (res *archive.PositionGetRes, err error) {
    res = new(archive.PositionGetRes)
	res.PositionInfoRes,err = service.Position().GetByPositionId(ctx, req.PositionId)
	return
}


// Add 添加职务
func (c *positionController) Add(ctx context.Context, req *archive.PositionAddReq) (res *archive.PositionAddRes, err error) {
	err = service.Position().Add(ctx, req)
	return
}


// Edit 修改职务
func (c *positionController) Edit(ctx context.Context, req *archive.PositionEditReq) (res *archive.PositionEditRes, err error) {
	err = service.Position().Edit(ctx, req)
	return
}


// Delete 删除职务
func (c *positionController) Delete(ctx context.Context, req *archive.PositionDeleteReq) (res *archive.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.PositionIds)
	return
}