// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-11-30 10:55:51
// 生成路径: api/v1/archive/t_education.go
// 生成人：xiao
// desc:学历相关参数
// company:云南奇讯科技有限公司
// ==========================================================================


package archive


import (    
    "github.com/gogf/gf/v2/frame/g"    
    commonApi "github.com/tiger1103/gfast/v3/api/v1/common"    
    "github.com/tiger1103/gfast/v3/internal/app/archive/model"
)
// EducationSearchReq 分页请求参数
type EducationSearchReq struct {
    g.Meta   `path:"/list" tags:"学历" method:"get" summary:"学历列表"`    
    EducationName  string `p:"educationName"` //学历    
    Status  string `p:"status" v:"status@integer#状态需为整数"` //状态    
    commonApi.PageReq
    commonApi.Author
}
// EducationSearchRes 列表返回结果
type EducationSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.EducationListRes `json:"list"`
}
// EducationAddReq 添加操作请求参数
type EducationAddReq struct {
    g.Meta   `path:"/add" tags:"学历" method:"post" summary:"学历添加"`
    commonApi.Author    
    EducationName  string   `p:"educationName" v:"required#学历不能为空"`    
    Seq  int   `p:"seq" `    
    Status  int   `p:"status" v:"required#状态不能为空"`    
    CreatedBy       uint64    
}
// EducationAddRes 添加操作返回结果
type EducationAddRes  struct {
    commonApi.EmptyRes
}
// EducationEditReq 修改操作请求参数
type EducationEditReq struct {
    g.Meta   `path:"/edit" tags:"学历" method:"put" summary:"学历修改"`
    commonApi.Author
    EducationId    int64  `p:"educationId" v:"required#主键ID不能为空"`    
    EducationName  string `p:"educationName" v:"required#学历不能为空"`    
    Seq  int `p:"seq" `    
    Status  int `p:"status" v:"required#状态不能为空"`    
    UpdatedBy       uint64    
}
// EducationEditRes 修改操作返回结果
type EducationEditRes  struct {
    commonApi.EmptyRes
}
// EducationGetReq 获取一条数据请求
type EducationGetReq struct {
    g.Meta `path:"/get" tags:"学历" method:"get" summary:"获取学历信息"`
    commonApi.Author
    EducationId    int64 `p:"educationId" v:"required#主键必须"` //通过主键获取
}
// EducationGetRes 获取一条数据结果
type EducationGetRes struct {
    g.Meta `mime:"application/json"`
    *model.EducationInfoRes
}
// EducationDeleteReq 删除数据请求
type EducationDeleteReq struct {
    g.Meta `path:"/delete" tags:"学历" method:"delete" summary:"删除学历"`
    commonApi.Author
    EducationIds    []int64 `p:"educationIds" v:"required#主键必须"` //通过主键删除
}
// EducationDeleteRes 删除数据返回
type EducationDeleteRes struct {
    commonApi.EmptyRes
}
