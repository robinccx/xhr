// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/model/entity/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================


package do


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// NurseChangeLog is the golang structure for table t_nurse_change_log.
type NurseChangeLog struct {
    gmeta.Meta   `orm:"table:t_nurse_change_log, do:true"`    
    Id       interface{}         `orm:"id,primary" json:"id"`    //    
    NurseId    interface{}         `orm:"nurse_id" json:"nurseId"`    //    
    ChangeType    interface{}         `orm:"change_type" json:"changeType"`    //    
    FromAreaId    interface{}         `orm:"from_area_id" json:"fromAreaId"`    //    
    FromDeptId    interface{}         `orm:"from_dept_id" json:"fromDeptId"`    //    
    FromEntityId    interface{}         `orm:"from_entity_id" json:"fromEntityId"`    //    
    FromEntityName    interface{}         `orm:"from_entity_name" json:"fromEntityName"`    //    
    ToAreaId    interface{}         `orm:"to_area_id" json:"toAreaId"`    //    
    ToDeptId    interface{}         `orm:"to_dept_id" json:"toDeptId"`    //    
    ToEntityId    interface{}         `orm:"to_entity_id" json:"toEntityId"`    //    
    ToEntityName    interface{}         `orm:"to_entity_name" json:"toEntityName"`    //    
    TransDate    *gtime.Time         `orm:"trans_date" json:"transDate"`    //    
    Status    interface{}         `orm:"status" json:"status"`    //    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    //    
    CreatedBy    interface{}         `orm:"created_by" json:"createdBy"`    //    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    interface{}         `orm:"updated_by" json:"updatedBy"`    //    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    interface{}         `orm:"deleted_by" json:"deletedBy"`    //    
}