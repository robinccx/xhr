// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: api/v1/archive/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询相关参数
// company:云南奇讯科技有限公司
// ==========================================================================


package archive


import (    
    "github.com/gogf/gf/v2/frame/g"    
    "github.com/gogf/gf/v2/os/gtime"    
    commonApi "github.com/tiger1103/gfast/v3/api/v1/common"    
    "github.com/tiger1103/gfast/v3/internal/app/archive/model"
)
// NurseChangeLogSearchReq 分页请求参数
type NurseChangeLogSearchReq struct {
    g.Meta   `path:"/list" tags:"变更记录查询" method:"get" summary:"变更记录查询列表"`    
    NurseId  string `p:"nurseId" v:"nurseId@integer#需为整数"` //    
    ChangeType  string `p:"changeType" v:"changeType@integer#需为整数"` //    
    FromAreaId  string `p:"fromAreaId" v:"fromAreaId@integer#需为整数"` //    
    FromDeptId  string `p:"fromDeptId" v:"fromDeptId@integer#需为整数"` //    
    ToAreaId  string `p:"toAreaId" v:"toAreaId@integer#需为整数"` //    
    ToDeptId  string `p:"toDeptId" v:"toDeptId@integer#需为整数"` //    
    TransDate  string `p:"transDate" v:"transDate@datetime#需为YYYY-MM-DD hh:mm:ss格式"` //    
    Status  string `p:"status" v:"status@integer#需为整数"` //    
    CreatedAt  string `p:"createdAt" v:"createdAt@datetime#需为YYYY-MM-DD hh:mm:ss格式"` //    
    CreatedBy  string `p:"createdBy"` //    
    DeletedBy  string `p:"deletedBy"` //    
    commonApi.PageReq
    commonApi.Author
}
// NurseChangeLogSearchRes 列表返回结果
type NurseChangeLogSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.NurseChangeLogListRes `json:"list"`
}
// NurseChangeLogAddReq 添加操作请求参数
type NurseChangeLogAddReq struct {
    g.Meta   `path:"/add" tags:"变更记录查询" method:"post" summary:"变更记录查询添加"`
    commonApi.Author    
    NurseId  int64   `p:"nurseId" `    
    ChangeType  int   `p:"changeType" `    
    FromAreaId  int   `p:"fromAreaId" `    
    FromDeptId  int64   `p:"fromDeptId" `    
    ToAreaId  int   `p:"toAreaId" `    
    ToDeptId  int64   `p:"toDeptId" `    
    TransDate  *gtime.Time   `p:"transDate" `    
    Status  int   `p:"status" v:"required#不能为空"`    
    DeletedBy  string   `p:"deletedBy" `    
    CreatedBy       uint64    
}
// NurseChangeLogAddRes 添加操作返回结果
type NurseChangeLogAddRes  struct {
    commonApi.EmptyRes
}
// NurseChangeLogEditReq 修改操作请求参数
type NurseChangeLogEditReq struct {
    g.Meta   `path:"/edit" tags:"变更记录查询" method:"put" summary:"变更记录查询修改"`
    commonApi.Author
    Id    int64  `p:"id" v:"required#主键ID不能为空"`    
    NurseId  int64 `p:"nurseId" `    
    ChangeType  int `p:"changeType" `    
    FromAreaId  int `p:"fromAreaId" `    
    FromDeptId  int64 `p:"fromDeptId" `    
    ToAreaId  int `p:"toAreaId" `    
    ToDeptId  int64 `p:"toDeptId" `    
    TransDate  *gtime.Time `p:"transDate" `    
    Status  int `p:"status" v:"required#不能为空"`    
    DeletedBy  string `p:"deletedBy" `    
    UpdatedBy       uint64    
}
// NurseChangeLogEditRes 修改操作返回结果
type NurseChangeLogEditRes  struct {
    commonApi.EmptyRes
}
// NurseChangeLogGetReq 获取一条数据请求
type NurseChangeLogGetReq struct {
    g.Meta `path:"/get" tags:"变更记录查询" method:"get" summary:"获取变更记录查询信息"`
    commonApi.Author
    Id    int64 `p:"id" v:"required#主键必须"` //通过主键获取
}
// NurseChangeLogGetRes 获取一条数据结果
type NurseChangeLogGetRes struct {
    g.Meta `mime:"application/json"`
    *model.NurseChangeLogInfoRes
}
// NurseChangeLogDeleteReq 删除数据请求
type NurseChangeLogDeleteReq struct {
    g.Meta `path:"/delete" tags:"变更记录查询" method:"delete" summary:"删除变更记录查询"`
    commonApi.Author
    Ids    []int64 `p:"ids" v:"required#主键必须"` //通过主键删除
}
// NurseChangeLogDeleteRes 删除数据返回
type NurseChangeLogDeleteRes struct {
    commonApi.EmptyRes
}
