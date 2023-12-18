// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/model/entity/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================


package entity


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Position is the golang structure for table t_position.
type Position struct {
    gmeta.Meta   `orm:"table:t_position"`    
    PositionId       int64         `orm:"position_id,primary" json:"positionId"`    // 职务ID    
    PositionCode    string         `orm:"position_code" json:"positionCode"`    // 职务代码    
    PositionName    string         `orm:"position_name" json:"positionName"`    // 职务    
    Seq    int         `orm:"seq" json:"seq"`    // 次序    
    Status    int         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    
    CreatedBy    string         `orm:"created_by" json:"createdBy"`    // 创建用户    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    string         `orm:"updated_by" json:"updatedBy"`    //    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    string         `orm:"deleted_by" json:"deletedBy"`    //
}
