// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-11-30 10:55:49
// 生成路径: api/v1/archive/t_title.go
// 生成人：xiao
// desc:职称相关参数
// company:云南奇讯科技有限公司
// ==========================================================================


package archive


import (    
    "github.com/gogf/gf/v2/frame/g"    
    commonApi "github.com/tiger1103/gfast/v3/api/v1/common"    
    "github.com/tiger1103/gfast/v3/internal/app/archive/model"
)
// TitleSearchReq 分页请求参数
type TitleSearchReq struct {
    g.Meta   `path:"/list" tags:"职称" method:"get" summary:"职称列表"`    
    TitleName  string `p:"titleName"` //联称    
    Status  string `p:"status" v:"status@integer#状态需为整数"` //状态    
    commonApi.PageReq
    commonApi.Author
}
// TitleSearchRes 列表返回结果
type TitleSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.TitleListRes `json:"list"`
}
// TitleAddReq 添加操作请求参数
type TitleAddReq struct {
    g.Meta   `path:"/add" tags:"职称" method:"post" summary:"职称添加"`
    commonApi.Author    
    TitleName  string   `p:"titleName" v:"required#联称不能为空"`    
    Seq  int   `p:"seq" `    
    Status  int   `p:"status" v:"required#状态不能为空"`    
    CreatedBy       uint64    
}
// TitleAddRes 添加操作返回结果
type TitleAddRes  struct {
    commonApi.EmptyRes
}
// TitleEditReq 修改操作请求参数
type TitleEditReq struct {
    g.Meta   `path:"/edit" tags:"职称" method:"put" summary:"职称修改"`
    commonApi.Author
    TitleId    int64  `p:"titleId" v:"required#主键ID不能为空"`    
    TitleName  string `p:"titleName" v:"required#联称不能为空"`    
    Seq  int `p:"seq" `    
    Status  int `p:"status" v:"required#状态不能为空"`    
    UpdatedBy       uint64    
}
// TitleEditRes 修改操作返回结果
type TitleEditRes  struct {
    commonApi.EmptyRes
}
// TitleGetReq 获取一条数据请求
type TitleGetReq struct {
    g.Meta `path:"/get" tags:"职称" method:"get" summary:"获取职称信息"`
    commonApi.Author
    TitleId    int64 `p:"titleId" v:"required#主键必须"` //通过主键获取
}
// TitleGetRes 获取一条数据结果
type TitleGetRes struct {
    g.Meta `mime:"application/json"`
    *model.TitleInfoRes
}
// TitleDeleteReq 删除数据请求
type TitleDeleteReq struct {
    g.Meta `path:"/delete" tags:"职称" method:"delete" summary:"删除职称"`
    commonApi.Author
    TitleIds    []int64 `p:"titleIds" v:"required#主键必须"` //通过主键删除
}
// TitleDeleteRes 删除数据返回
type TitleDeleteRes struct {
    commonApi.EmptyRes
}
