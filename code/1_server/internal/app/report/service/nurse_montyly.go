package service

import (
	"context"
	"github.com/tiger1103/gfast/v3/api/v1/report"
)

type INurseMonthly interface {
	ListTotal(ctx context.Context, req *report.MonthlyTotalSearchReq) (res *report.MonthlyTotalSearchRes, err error)
	ListTitle(ctx context.Context, req *report.MonthlyTitleSearchReq) (res *report.MonthlyTitleSearchRes, err error)
	ListEdu(ctx context.Context, req *report.MonthlyEduSearchReq) (res *report.MonthlyEduSearchRes, err error)
	ListQuarterTotal(ctx context.Context, req *report.QuarterTotalSearchReq) (res *report.QuarterTotalSearchRes, err error)
	GetHomeData(ctx context.Context, req *report.HomeSearchReq) (res *report.HomeSearchRes, err error)
}

var localNurseMonthly INurseMonthly

func NurseMonthly() INurseMonthly {
	if localNurseMonthly == nil {
		panic("implement not found for interface INurseMonthly, forgot register?")
	}
	return localNurseMonthly
}

func RegisterNurseMonthly(i INurseMonthly) {
	localNurseMonthly = i
}
