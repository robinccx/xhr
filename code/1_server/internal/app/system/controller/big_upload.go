package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/common/service"
)

var BigUpload = new(bigUploadController)

type bigUploadController struct{}

// 上传分片文件
func (c *bigUploadController) Upload(ctx context.Context, req *system.BigUploadReq) (res *system.BigUploadRes, err error) {
	r := g.RequestFromCtx(ctx)
	req = new(system.BigUploadReq)
	err = req.Bind(r.Request)
	if err != nil {
		return
	}
	return service.BigUpload().Upload(ctx, req)
}

// 上传文件检查
func (c *bigUploadController) UploadCheck(ctx context.Context, req *system.BigUploadCheckReq) (res *system.BigUploadCheckRes, err error) {
	r := g.RequestFromCtx(ctx)
	req = new(system.BigUploadCheckReq)
	err = req.Bind(r.Request)
	if err != nil {
		return
	}
	return service.BigUpload().UploadCheck(ctx, req)
}

// 上传文件合并
func (c *bigUploadController) UploadMerge(ctx context.Context, req *system.BigUploadMergeReq) (res *system.BigUploadRes, err error) {
	r := g.RequestFromCtx(ctx)
	req = new(system.BigUploadMergeReq)
	err = req.Bind(r.Request)
	if err != nil {
		return
	}
	return service.BigUpload().UploadMerge(ctx, req)
}