// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/model/entity/t_department.go
// 生成人：xiao
// desc:科室
// company:云南奇讯科技有限公司
// ==========================================================================


package do


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Department is the golang structure for table t_department.
type Department struct {
    gmeta.Meta   `orm:"table:t_department, do:true"`    
    DeptId       interface{}         `orm:"dept_id,primary" json:"deptId"`    // 科室ID    
    ParentId    interface{}         `orm:"parent_id" json:"parentId"`    // 父级ID    
    DeptCode    interface{}         `orm:"dept_code" json:"deptCode"`    // 科室代码    
    DeptName    interface{}         `orm:"dept_name" json:"deptName"`    // 科室名称    
    DeptType    interface{}         `orm:"dept_type" json:"deptType"`    // 片区    
    Leader    interface{}         `orm:"leader" json:"leader"`    // 负责人    
    Phone    interface{}         `orm:"phone" json:"phone"`    // 电话    
    Email    interface{}         `orm:"email" json:"email"`    // 邮件    
    Att1    interface{}         `orm:"att1" json:"att1"`    // 属性1    
    Att2    interface{}         `orm:"att2" json:"att2"`    // 属性2    
    Att3    interface{}         `orm:"att3" json:"att3"`    // 属性3    
    Status    interface{}         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    
    CreatedBy    interface{}         `orm:"created_by" json:"createdBy"`    // 创建用户    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    interface{}         `orm:"updated_by" json:"updatedBy"`    // 修改用户    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    interface{}         `orm:"deleted_by" json:"deletedBy"`    // 删除用户    
}