package controller

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/report"
	"github.com/tiger1103/gfast/v3/internal/app/report/service"
	systemController "github.com/tiger1103/gfast/v3/internal/app/system/controller"
)

type nurseMonthlyController struct {
	systemController.BaseController
}

var NurseMonthly = new(nurseMonthlyController)

// 月度入职统计
func (c *nurseMonthlyController) ListTotal(ctx context.Context, req *report.MonthlyTotalSearchReq) (res *report.MonthlyTotalSearchRes, err error) {
	res, err = service.NurseMonthly().ListTotal(ctx, req)
	return
}

// 月度职称统计
func (c *nurseMonthlyController) ListTitle(ctx context.Context, req *report.MonthlyTitleSearchReq) (res *report.MonthlyTitleSearchRes, err error) {
	res, err = service.NurseMonthly().ListTitle(ctx, req)
	return
}

// 月度学历统计
func (c *nurseMonthlyController) ListEdu(ctx context.Context, req *report.MonthlyEduSearchReq) (res *report.MonthlyEduSearchRes, err error) {
	res, err = service.NurseMonthly().ListEdu(ctx, req)
	return
}

// 月度学历统计
func (c *nurseMonthlyController) ListQuarterTotal(ctx context.Context, req *report.QuarterTotalSearchReq) (res *report.QuarterTotalSearchRes, err error) {
	res, err = service.NurseMonthly().ListQuarterTotal(ctx, req)
	return
}

// 月度学历统计
func (c *nurseMonthlyController) GetHomeData(ctx context.Context, req *report.HomeSearchReq) (res *report.HomeSearchRes, err error) {
	res, err = service.NurseMonthly().GetHomeData(ctx, req)
	return
}
