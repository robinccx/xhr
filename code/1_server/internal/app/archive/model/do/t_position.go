// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/model/entity/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================


package do


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Position is the golang structure for table t_position.
type Position struct {
    gmeta.Meta   `orm:"table:t_position, do:true"`    
    PositionId       interface{}         `orm:"position_id,primary" json:"positionId"`    // 职务ID    
    PositionCode    interface{}         `orm:"position_code" json:"positionCode"`    // 职务代码    
    PositionName    interface{}         `orm:"position_name" json:"positionName"`    // 职务    
    Seq    interface{}         `orm:"seq" json:"seq"`    // 次序    
    Status    interface{}         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    
    CreatedBy    interface{}         `orm:"created_by" json:"createdBy"`    // 创建用户    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    interface{}         `orm:"updated_by" json:"updatedBy"`    //    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    interface{}         `orm:"deleted_by" json:"deletedBy"`    //    
}