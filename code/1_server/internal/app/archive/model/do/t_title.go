// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-11-30 18:32:51
// 生成路径: internal/app/archive/model/entity/t_title.go
// 生成人：xiao
// desc:职称
// company:云南奇讯科技有限公司
// ==========================================================================


package do


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Title is the golang structure for table t_title.
type Title struct {
    gmeta.Meta   `orm:"table:t_title, do:true"`    
    TitleId       interface{}         `orm:"title_id,primary" json:"titleId"`    // ID    
    TitleName    interface{}         `orm:"title_name" json:"titleName"`    // 职称    
    Seq    interface{}         `orm:"seq" json:"seq"`    // 次序    
    Status    interface{}         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    
    CreatedBy    interface{}         `orm:"created_by" json:"createdBy"`    // 创建用户    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    interface{}         `orm:"updated_by" json:"updatedBy"`    // 修改用户    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    interface{}         `orm:"deleted_by" json:"deletedBy"`    // 删除用户    
}