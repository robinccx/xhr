// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-16 10:46:43
// 生成路径: internal/app/archive/model/entity/t_nurse_import.go
// 生成人：xiao
// desc:护士导入
// company:云南奇讯科技有限公司
// ==========================================================================


package entity


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// NurseImport is the golang structure for table t_nurse_import.
type NurseImport struct {
    gmeta.Meta   `orm:"table:t_nurse_import"`    
    Id       int64         `orm:"id,primary" json:"id"`    //    
    SessionId    string         `orm:"session_id" json:"sessionId"`    //    
    NurseId    int64         `orm:"nurse_id" json:"nurseId"`    //    
    NurseCode    string         `orm:"nurse_code" json:"nurseCode"`    //    
    NurseName    string         `orm:"nurse_name" json:"nurseName"`    //    
    StartDate    string         `orm:"start_date" json:"startDate"`    //    
    EndDate    string         `orm:"end_date" json:"endDate"`    //    
    Sex    string         `orm:"sex" json:"sex"`    //    
    Birthday    string         `orm:"birthday" json:"birthday"`    //    
    AreaId    int         `orm:"area_id" json:"areaId"`    //    
    AreaCode    string         `orm:"area_code" json:"areaCode"`    //    
    DeptId    int64         `orm:"dept_id" json:"deptId"`    //    
    DeptCode    string         `orm:"dept_code" json:"deptCode"`    //    
    EducationId    int64         `orm:"education_id" json:"educationId"`    //    
    EducationCode    string         `orm:"education_code" json:"educationCode"`    //    
    TitleId    int64         `orm:"title_id" json:"titleId"`    //    
    TitleCode    string         `orm:"title_code" json:"titleCode"`    //    
    StaffType    int         `orm:"staff_type" json:"staffType"`    //    
    StaffTypeCode    string         `orm:"staff_type_code" json:"staffTypeCode"`    //    
    CertNo    string         `orm:"cert_no" json:"certNo"`    //    
    CertEnddate    *gtime.Time         `orm:"cert_end_date" json:"certEnddate"`    //    
    RegDate    *gtime.Time         `orm:"reg_date" json:"regDate"`    //    
    Note    string         `orm:"note" json:"note"`    //    
    Seq    int         `orm:"seq" json:"seq"`    //    
    Status    int         `orm:"status" json:"status"`    //    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    //    
    CreatedBy    string         `orm:"created_by" json:"createdBy"`    //
}
