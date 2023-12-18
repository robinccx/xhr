// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-16 10:46:42
// 生成路径: internal/app/archive/model/entity/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================


package do


import (    
    "github.com/gogf/gf/v2/os/gtime"    
    "github.com/gogf/gf/v2/util/gmeta"
)


// Nurse is the golang structure for table t_nurse.
type Nurse struct {
    gmeta.Meta   `orm:"table:t_nurse, do:true"`    
    NurseId       interface{}         `orm:"nurse_id,primary" json:"nurseId"`    // ID    
    NurseCode    interface{}         `orm:"nurse_code" json:"nurseCode"`    // 编号    
    NurseName    interface{}         `orm:"nurse_name" json:"nurseName"`    // 名称    
    Sex    interface{}         `orm:"sex" json:"sex"`    // 性别    
    Birthday    *gtime.Time         `orm:"birthday" json:"birthday"`    // 出生年月    
    AreaId    interface{}         `orm:"area_id" json:"areaId"`    // 院区    
    DeptId    interface{}         `orm:"dept_id" json:"deptId"`    // 科室    
    StartDate    *gtime.Time         `orm:"start_date" json:"startDate"`    // 入职日期    
    EducationId    interface{}         `orm:"education_id" json:"educationId"`    // 学历    
    TitleId    interface{}         `orm:"title_id" json:"titleId"`    // 职称    
    StaffType    interface{}         `orm:"staff_type" json:"staffType"`    // 人员类别    
    Note    interface{}         `orm:"note" json:"note"`    // 备注    
    EndDate    *gtime.Time         `orm:"end_date" json:"endDate"`    // 离职日期    
    Status    interface{}         `orm:"status" json:"status"`    // 状态    
    CreatedAt    *gtime.Time         `orm:"created_at" json:"createdAt"`    // 创建时间    
    CreatedBy    interface{}         `orm:"created_by" json:"createdBy"`    // 创建用户    
    UpdatedAt    *gtime.Time         `orm:"updated_at" json:"updatedAt"`    //    
    UpdatedBy    interface{}         `orm:"updated_by" json:"updatedBy"`    // 修改用户    
    DeletedAt    *gtime.Time         `orm:"deleted_at" json:"deletedAt"`    //    
    DeletedBy    interface{}         `orm:"deleted_by" json:"deletedBy"`    // 删除用户    
    SessionId    interface{}         `orm:"session_id" json:"sessionId"`    //    
    CertEnddate    *gtime.Time         `orm:"cert_end_date" json:"certEnddate"`    // 职业证书到期日期    
    RegDate    *gtime.Time         `orm:"reg_date" json:"regDate"`    // 注册或变更日期    
    CertNo    interface{}         `orm:"cert_no" json:"certNo"`    // 职业证书    
    IdNo    interface{}         `orm:"id_no" json:"idNo"`    // 身份证号    
    PositionId    interface{}         `orm:"position_id" json:"positionId"`    // 职务    
}