// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-12-01 20:01:44
// 生成路径: api/v1/archive/t_nurse.go
// 生成人：xiao
// desc:护士相关参数
// company:云南奇讯科技有限公司
// ==========================================================================

package archive

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	commonApi "github.com/tiger1103/gfast/v3/api/v1/common"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"
)

// NurseSearchReq 分页请求参数
type NurseSearchReq struct {
	g.Meta      `path:"/list" tags:"护士" method:"get" summary:"护士列表"`
	AreaId      string `p:"areaId" v:"areaId@integer#院区需为整数"`                             //院区
	NurseName   string `p:"nurseName"`                                                    //名称
	DeptId      string `p:"deptId" v:"deptId@integer#科室需为整数"`                             //科室
	Status      string `p:"status" v:"status@integer#状态需为整数"`                             //状态
	RegStatus   string `p:"regStatus" v:"regStatus@integer#状态需为整数"`                       //注册与变更状态
	StartDate   string `p:"startDate" v:"startDate@datetime#入职日期需为YYYY-MM-DD hh:mm:ss格式"` //入职日期
	EducationId string `p:"educationId" v:"educationId@integer#学历需为整数"`                   //学历
	TitleId     string `p:"titleId" v:"titleId@integer#职称需为整数"`                           //职称
	StaffType   string `p:"staffType" v:"staffType@integer#人员类别需为整数"`                     //人员类别
	PositionId  string `p:"positionId" v:"positionId@integer#职务需为整数"`                     //职称
	NurseCode   string `p:"nurseCode"`                                                    //编号
	commonApi.PageReq
	commonApi.Author
}

// NurseSearchRes 列表返回结果
type NurseSearchRes struct {
	g.Meta `mime:"application/json"`
	commonApi.ListRes
	List []*model.NurseListRes `json:"list"`
}

// NurseAddReq 添加操作请求参数
type NurseAddReq struct {
	g.Meta `path:"/add" tags:"护士" method:"post" summary:"护士添加"`
	commonApi.Author
	NurseCode   string      `p:"nurseCode" `
	NurseName   string      `p:"nurseName" v:"required#名称不能为空"`
	Sex         int         `p:"sex" `
	Birthday    *gtime.Time `p:"birthday" `
	AreaId      int         `p:"areaId" `
	DeptId      int64       `p:"deptId" `
	StartDate   *gtime.Time `p:"startDate" `
	EducationId int64       `p:"educationId" `
	TitleId     int64       `p:"titleId" `
	StaffType   int         `p:"staffType" `
	Note        string      `p:"note" `
	EndDate     *gtime.Time `p:"endDate" `
	CertNo      string      `p:"certNo"`      // 职业证书
	CertEnddate *gtime.Time `p:"certEndDate"` // 职业证书到期日期
	RegDate     *gtime.Time `p:"regDate"`     // 注册或变更日期
	Status      int         `p:"status" v:"required#状态不能为空"`
	CreatedBy   uint64
}

// NurseAddRes 添加操作返回结果
type NurseAddRes struct {
	commonApi.EmptyRes
}

// NurseEditReq 修改操作请求参数
type NurseEditReq struct {
	g.Meta `path:"/edit" tags:"护士" method:"put" summary:"护士修改"`
	commonApi.Author
	NurseId     int64       `p:"nurseId" v:"required#主键ID不能为空"`
	NurseCode   string      `p:"nurseCode" `
	NurseName   string      `p:"nurseName" v:"required#名称不能为空"`
	Sex         int         `p:"sex" `
	Birthday    *gtime.Time `p:"birthday" `
	AreaId      int         `p:"areaId" `
	DeptId      int64       `p:"deptId" `
	StartDate   *gtime.Time `p:"startDate" `
	EducationId int64       `p:"educationId" `
	PositionId  int64       `p:"positionId" `
	TitleId     int64       `p:"titleId" `
	StaffType   int         `p:"staffType" `
	Note        string      `p:"note" `
	EndDate     *gtime.Time `p:"endDate" `
	IdNo        string      `p:"idNo"`
	CertNo      string      `p:"certNo"`      // 职业证书
	CertEndDate *gtime.Time `p:"certEndDate"` // 职业证书到期日期
	RegDate     *gtime.Time `p:"regDate"`     // 注册或变更日期
	Status      int         `p:"status" v:"required#状态不能为空"`
	UpdatedBy   uint64
}

// NurseEditRes 修改操作返回结果
type NurseEditRes struct {
	commonApi.EmptyRes
}

// NurseGetReq 获取一条数据请求
type NurseGetReq struct {
	g.Meta `path:"/get" tags:"护士" method:"get" summary:"获取护士信息"`
	commonApi.Author
	NurseId int64 `p:"nurseId" v:"required#主键必须"` //通过主键获取
}

// NurseGetRes 获取一条数据结果
type NurseGetRes struct {
	g.Meta `mime:"application/json"`
	*model.NurseInfoRes
}

// NurseDeleteReq 删除数据请求
type NurseDeleteReq struct {
	g.Meta `path:"/delete" tags:"护士" method:"delete" summary:"删除护士"`
	commonApi.Author
	NurseIds []int64 `p:"nurseIds" v:"required#主键必须"` //通过主键删除
}

// NurseDeleteRes 删除数据返回
type NurseDeleteRes struct {
	commonApi.EmptyRes
}

// NurseImportReq 导入请求参数
type NurseImportReq struct {
	g.Meta `path:"/import" tags:"导入" method:"post" summary:"导入"`
	AreaId int `json:"areaId"` // 院区
	Data   [][]string
	commonApi.Author
}

// NurseImportRes 导入返回结果
type NurseImportRes struct {
	SuccessCount int                      `json:"successCount"` //
	ExistsCount  int                      `json:"existsCount"`  //
	ErrorCount   int                      `json:"errorCount"`   //
	ErrorData    []*model.ImportErrorItem `json:"errorData"`    //
}

// NurseUpdateByImportReq 导入请求参数
type NurseUpdateByImportReq struct {
	g.Meta `path:"/updatebyimport" tags:"导入" method:"post" summary:"导入"`
	AreaId int `json:"areaId"` // 院区
	Data   [][]string
	commonApi.Author
}
