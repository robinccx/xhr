// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-16 10:46:43
// 生成路径: internal/app/archive/model/entity/t_nurse_import.go
// 生成人：xiao
// desc:护士导入
// company:云南奇讯科技有限公司
// ==========================================================================


package do


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// NurseImport is the golang structure for table t_nurse_import.
type NurseImport struct {
    gmeta.Meta   `orm:"table:t_nurse_import, do:true"`    
    Id       interface{}         `orm:"id,primary" json:"id"`    //    
    SessionId    interface{}         `orm:"session_id" json:"sessionId"`    //    
    NurseId    interface{}         `orm:"nurse_id" json:"nurseId"`    //    
    NurseCode    interface{}         `orm:"nurse_code" json:"nurseCode"`    //    
    NurseName    interface{}         `orm:"nurse_name" json:"nurseName"`    //    
    StartDate    interface{}         `orm:"start_date" json:"startDate"`    //    
    EndDate    interface{}         `orm:"end_date" json:"endDate"`    //    
    Sex    interface{}         `orm:"sex" json:"sex"`    //    
    Birthday    interface{}         `orm:"birthday" json:"birthday"`    //    
    AreaId    interface{}         `orm:"area_id" json:"areaId"`    //    
    AreaCode    interface{}         `orm:"area_code" json:"areaCode"`    //    
    DeptId    interface{}         `orm:"dept_id" json:"deptId"`    //    
    DeptCode    interface{}         `orm:"dept_code" json:"deptCode"`    //    
    EducationId    interface{}         `orm:"education_id" json:"educationId"`    //    
    EducationCode    interface{}         `orm:"education_code" json:"educationCode"`    //    
    TitleId    interface{}         `orm:"title_id" json:"titleId"`    //    
    TitleCode    interface{}         `orm:"title_code" json:"titleCode"`    //    
    StaffType    interface{}         `orm:"staff_type" json:"staffType"`    //    
    StaffTypeCode    interface{}         `orm:"staff_type_code" json:"staffTypeCode"`    //    
    CertNo    interface{}         `orm:"cert_no" json:"certNo"`    //    
    CertEnddate    *gtime.Time         `orm:"cert_end_date" json:"certEnddate"`    //    
    RegDate    *gtime.Time         `orm:"reg_date" json:"regDate"`    //    
    Note    interface{}         `orm:"note" json:"note"`    //    
    Seq    interface{}         `orm:"seq" json:"seq"`    //    
    Status    interface{}         `orm:"status" json:"status"`    //    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    //    
    CreatedBy    interface{}         `orm:"created_by" json:"createdBy"`    //    
}