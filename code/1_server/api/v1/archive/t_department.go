// ==========================================================================
// GFast自动生成api操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: api/v1/archive/t_department.go
// 生成人：xiao
// desc:科室相关参数
// company:云南奇讯科技有限公司
// ==========================================================================


package archive


import (    
    "github.com/gogf/gf/v2/frame/g"    
    commonApi "github.com/tiger1103/gfast/v3/api/v1/common"    
    "github.com/tiger1103/gfast/v3/internal/app/archive/model"
)
// DepartmentSearchReq 分页请求参数
type DepartmentSearchReq struct {
    g.Meta   `path:"/list" tags:"科室" method:"get" summary:"科室列表"`    
    DeptCode  string `p:"deptCode"` //科室代码    
    DeptName  string `p:"deptName"` //科室名称    
    Status  string `p:"status" v:"status@integer#状态需为整数"` //状态    
    DeptType  string `p:"deptType"` //片区    
    Att1  string `p:"att1"` //属性1    
    Att2  string `p:"att2"` //属性2    
    Att3  string `p:"att3"` //属性3    
    commonApi.PageReq
    commonApi.Author
}
// DepartmentSearchRes 列表返回结果
type DepartmentSearchRes struct {
    g.Meta `mime:"application/json"`
    commonApi.ListRes
    List []*model.DepartmentListRes `json:"list"`
}
// DepartmentAddReq 添加操作请求参数
type DepartmentAddReq struct {
    g.Meta   `path:"/add" tags:"科室" method:"post" summary:"科室添加"`
    commonApi.Author    
    DeptCode  string   `p:"deptCode" v:"required#科室代码不能为空"`    
    DeptName  string   `p:"deptName" v:"required#科室名称不能为空"`    
    DeptType  string   `p:"deptType" `    
    Leader  string   `p:"leader" `    
    Phone  string   `p:"phone" `    
    Email  string   `p:"email" `    
    Att1  string   `p:"att1" `    
    Att2  string   `p:"att2" `    
    Att3  string   `p:"att3" `    
    Status  int   `p:"status" v:"required#状态不能为空"`    
    CreatedBy       uint64    
}
// DepartmentAddRes 添加操作返回结果
type DepartmentAddRes  struct {
    commonApi.EmptyRes
}
// DepartmentEditReq 修改操作请求参数
type DepartmentEditReq struct {
    g.Meta   `path:"/edit" tags:"科室" method:"put" summary:"科室修改"`
    commonApi.Author
    DeptId    int64  `p:"deptId" v:"required#主键ID不能为空"`    
    DeptCode  string `p:"deptCode" v:"required#科室代码不能为空"`    
    DeptName  string `p:"deptName" v:"required#科室名称不能为空"`    
    DeptType  string `p:"deptType" `    
    Leader  string `p:"leader" `    
    Phone  string `p:"phone" `    
    Email  string `p:"email" `    
    Att1  string `p:"att1" `    
    Att2  string `p:"att2" `    
    Att3  string `p:"att3" `    
    Status  int `p:"status" v:"required#状态不能为空"`    
    UpdatedBy       uint64    
}
// DepartmentEditRes 修改操作返回结果
type DepartmentEditRes  struct {
    commonApi.EmptyRes
}
// DepartmentGetReq 获取一条数据请求
type DepartmentGetReq struct {
    g.Meta `path:"/get" tags:"科室" method:"get" summary:"获取科室信息"`
    commonApi.Author
    DeptId    int64 `p:"deptId" v:"required#主键必须"` //通过主键获取
}
// DepartmentGetRes 获取一条数据结果
type DepartmentGetRes struct {
    g.Meta `mime:"application/json"`
    *model.DepartmentInfoRes
}
// DepartmentDeleteReq 删除数据请求
type DepartmentDeleteReq struct {
    g.Meta `path:"/delete" tags:"科室" method:"delete" summary:"删除科室"`
    commonApi.Author
    DeptIds    []int64 `p:"deptIds" v:"required#主键必须"` //通过主键删除
}
// DepartmentDeleteRes 删除数据返回
type DepartmentDeleteRes struct {
    commonApi.EmptyRes
}
