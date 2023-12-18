package report

import (
	"github.com/gogf/gf/v2/frame/g"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/report/model"
)

// MonthlyTotalSearchReq 分页请求参数
type MonthlyTotalSearchReq struct {
	g.Meta    `path:"/monthlytotal" tags:"月度入职统计" method:"get" summary:"月度入职统计"`
	AreaId    string `p:"area_id"`                                //院区
	Year      string `p:"year"`                                   //年份
	RegStatus string `p:"regStatus" v:"regStatus@integer#状态需为整数"` //注册与变更状态
	commonApi.PageReq
	commonApi.Author
}

// MonthlyTotalSearchRes 列表返回结果
type MonthlyTotalSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.MonthlyTotalInfoRes `json:"list"`
}

// MonthlyTitleSearchReq 分页请求参数
type MonthlyTitleSearchReq struct {
	g.Meta    `path:"/monthlytitle" tags:"月度职称统计" method:"get" summary:"月度职称统计"`
	AreaId    string `p:"area_id"`                                //院区
	Year      string `p:"year"`                                   //年份
	RegStatus string `p:"regStatus" v:"regStatus@integer#状态需为整数"` //注册与变更状态
	commonApi.PageReq
	commonApi.Author
}

// type MonthlyTitleSearchRes struct { 列表返回结果
type MonthlyTitleSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []map[string]interface{} `json:"list"`
}

// MonthlyEduSearchReq 分页请求参数
type MonthlyEduSearchReq struct {
	g.Meta    `path:"/monthlyedu" tags:"月度学历统计" method:"get" summary:"月度学历统计"`
	AreaId    string `p:"area_id"`                                //院区
	Year      string `p:"year"`                                   //年份
	RegStatus string `p:"regStatus" v:"regStatus@integer#状态需为整数"` //注册与变更状态
	commonApi.PageReq
	commonApi.Author
}

// type MonthlyEduSearchRes struct { 列表返回结果
type MonthlyEduSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []map[string]interface{} `json:"list"`
}

// QuarterTotalSearchReq 分页请求参数
type QuarterTotalSearchReq struct {
	g.Meta `path:"/quartertotal" tags:"月度学历统计" method:"get" summary:"月度学历统计"`
	Year   string `p:"year"` //年份
	commonApi.PageReq
	commonApi.Author
}

// type QuarterTotalSearchRes struct { 列表返回结果
type QuarterTotalSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.QuarterTotalInfoRes `json:"list"`
}

func (q *QuarterTotalSearchRes) AppendItem(year string, itemClass string, itemName string, itemKey string) {
	q.List = append(q.List, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  itemName,
		ItemKey:   itemKey,
	})
}

// HomeSearchReq 分页请求参数
type HomeSearchReq struct {
	g.Meta `path:"/home" tags:"获取首页数据" method:"get" summary:"获取首页数据"`
	commonApi.PageReq
	commonApi.Author
}

// HomeSearchReq 总在职数
type HomeSearchTotalRes struct {
	ItemName   string `json:"itemName"`
	ItemValue  string `json:"itemValue"`
	ItemValue2 string `json:"itemValue2"`
}

// HomeSearchReq 按院区，每月在职数
type HomeSearchAreaGroupRes struct {
	AreaList  []string `json:"areaList"`
	Data      [][]int  `json:"data"`
	LeaveData [][]int  `json:"leaveData"`
}

// HomeSearchEduGroupRes 按学历汇总，在职数
type HomeSearchEduGroupRes struct {
	ItemName  string `json:"itemName"`
	ItemValue int    `json:"itemValue"`
}

// HomeSearchReq 首面返回信息
type HomeSearchRes struct {
	TotalData []*HomeSearchTotalRes   `json:"total"`
	AreaData  HomeSearchAreaGroupRes  `json:"area"`
	EduData   []HomeSearchEduGroupRes `json:"edu"`
	TitleData []HomeSearchEduGroupRes `json:"title"`
}
