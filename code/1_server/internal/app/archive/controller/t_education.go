// ==========================================================================
// GFast自动生成controller操作代码。
// 生成日期：2023-11-30 18:32:48
// 生成路径: internal/app/archive/controller/t_education.go
// 生成人：xiao
// desc:学历
// company:云南奇讯科技有限公司
// ==========================================================================


package controller


import (
    "context"
    "github.com/tiger1103/gfast/v3/api/v1/archive"
    "github.com/tiger1103/gfast/v3/internal/app/archive/service"    
    systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"    
)


type educationController struct {    
    systemController.BaseController    
}


var Education = new(educationController)


// List 列表
func (c *educationController) List(ctx context.Context, req *archive.EducationSearchReq) (res *archive.EducationSearchRes, err error) {
	res, err = service.Education().List(ctx, req)
	return
}


// Get 获取学历
func (c *educationController) Get(ctx context.Context, req *archive.EducationGetReq) (res *archive.EducationGetRes, err error) {
    res = new(archive.EducationGetRes)
	res.EducationInfoRes,err = service.Education().GetByEducationId(ctx, req.EducationId)
	return
}


// Add 添加学历
func (c *educationController) Add(ctx context.Context, req *archive.EducationAddReq) (res *archive.EducationAddRes, err error) {
	err = service.Education().Add(ctx, req)
	return
}


// Edit 修改学历
func (c *educationController) Edit(ctx context.Context, req *archive.EducationEditReq) (res *archive.EducationEditRes, err error) {
	err = service.Education().Edit(ctx, req)
	return
}


// Delete 删除学历
func (c *educationController) Delete(ctx context.Context, req *archive.EducationDeleteReq) (res *archive.EducationDeleteRes, err error) {
	err = service.Education().Delete(ctx, req.EducationIds)
	return
}