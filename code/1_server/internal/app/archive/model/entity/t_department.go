// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/model/entity/t_department.go
// 生成人：xiao
// desc:科室
// company:云南奇讯科技有限公司
// ==========================================================================


package entity


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Department is the golang structure for table t_department.
type Department struct {
    gmeta.Meta   `orm:"table:t_department"`    
    DeptId       int64         `orm:"dept_id,primary" json:"deptId"`    // 科室ID    
    ParentId    int64         `orm:"parent_id" json:"parentId"`    // 父级ID    
    DeptCode    string         `orm:"dept_code" json:"deptCode"`    // 科室代码    
    DeptName    string         `orm:"dept_name" json:"deptName"`    // 科室名称    
    DeptType    string         `orm:"dept_type" json:"deptType"`    // 片区    
    Leader    string         `orm:"leader" json:"leader"`    // 负责人    
    Phone    string         `orm:"phone" json:"phone"`    // 电话    
    Email    string         `orm:"email" json:"email"`    // 邮件    
    Att1    string         `orm:"att1" json:"att1"`    // 属性1    
    Att2    string         `orm:"att2" json:"att2"`    // 属性2    
    Att3    string         `orm:"att3" json:"att3"`    // 属性3    
    Status    int         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    
    CreatedBy    string         `orm:"created_by" json:"createdBy"`    // 创建用户    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    string         `orm:"updated_by" json:"updatedBy"`    // 修改用户    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    string         `orm:"deleted_by" json:"deletedBy"`    // 删除用户
}
