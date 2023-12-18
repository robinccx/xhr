// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-12-16 10:37:06
// 生成路径: api/v1/archive/t_position.go
// 生成人：xiao
// desc:职务相关参数
// company:云南奇讯科技有限公司
// ==========================================================================


package archive


import (    
    "github.com/gogf/gf/v2/frame/g"    
    commonApi "github.com/tiger1103/gfast/v3/api/v1/common"    
    "github.com/tiger1103/gfast/v3/internal/app/archive/model"
)
// PositionSearchReq 分页请求参数
type PositionSearchReq struct {
    g.Meta   `path:"/list" tags:"职务" method:"get" summary:"职务列表"`    
    PositionName  string `p:"positionName"` //职务    
    commonApi.PageReq
    commonApi.Author
}
// PositionSearchRes 列表返回结果
type PositionSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.PositionListRes `json:"list"`
}
// PositionAddReq 添加操作请求参数
type PositionAddReq struct {
    g.Meta   `path:"/add" tags:"职务" method:"post" summary:"职务添加"`
    commonApi.Author    
    PositionName  string   `p:"positionName" v:"required#职务不能为空"`    
    Seq  int   `p:"seq" `    
    Status  int   `p:"status" v:"required#状态不能为空"`    
    CreatedBy       uint64    
}
// PositionAddRes 添加操作返回结果
type PositionAddRes  struct {
    commonApi.EmptyRes
}
// PositionEditReq 修改操作请求参数
type PositionEditReq struct {
    g.Meta   `path:"/edit" tags:"职务" method:"put" summary:"职务修改"`
    commonApi.Author
    PositionId    int64  `p:"positionId" v:"required#主键ID不能为空"`    
    PositionName  string `p:"positionName" v:"required#职务不能为空"`    
    Seq  int `p:"seq" `    
    Status  int `p:"status" v:"required#状态不能为空"`    
    UpdatedBy       uint64    
}
// PositionEditRes 修改操作返回结果
type PositionEditRes  struct {
    commonApi.EmptyRes
}
// PositionGetReq 获取一条数据请求
type PositionGetReq struct {
    g.Meta `path:"/get" tags:"职务" method:"get" summary:"获取职务信息"`
    commonApi.Author
    PositionId    int64 `p:"positionId" v:"required#主键必须"` //通过主键获取
}
// PositionGetRes 获取一条数据结果
type PositionGetRes struct {
    g.Meta `mime:"application/json"`
    *model.PositionInfoRes
}
// PositionDeleteReq 删除数据请求
type PositionDeleteReq struct {
    g.Meta `path:"/delete" tags:"职务" method:"delete" summary:"删除职务"`
    commonApi.Author
    PositionIds    []int64 `p:"positionIds" v:"required#主键必须"` //通过主键删除
}
// PositionDeleteRes 删除数据返回
type PositionDeleteRes struct {
    commonApi.EmptyRes
}
