// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-11-30 18:32:48
// 生成路径: internal/app/archive/model/entity/t_education.go
// 生成人：xiao
// desc:学历
// company:云南奇讯科技有限公司
// ==========================================================================


package entity


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Education is the golang structure for table t_education.
type Education struct {
    gmeta.Meta   `orm:"table:t_education"`    
    EducationId       int64         `orm:"education_id,primary" json:"educationId"`    // ID    
    EducationName    string         `orm:"education_name" json:"educationName"`    // 学历    
    Seq    int         `orm:"seq" json:"seq"`    // 次序    
    Status    int         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    //    
    CreatedBy    string         `orm:"created_by" json:"createdBy"`    //    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    string         `orm:"updated_by" json:"updatedBy"`    //    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    string         `orm:"deleted_by" json:"deletedBy"`    //
}
