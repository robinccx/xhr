// ==========================================================================
// GFast自动生成model entity操作代码。
// 生成日期：2023-12-16 10:46:42
// 生成路径: internal/app/archive/model/entity/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================

package entity

import (
    "github.com/gogf/gf/v2/os/gtime"
    "github.com/gogf/gf/v2/util/gmeta"
)

// Nurse is the golang structure for table t_nurse.
type Nurse struct {
    gmeta.Meta        `orm:"table:t_nurse"`
    NurseId           int64                  `orm:"nurse_id,primary" json:"nurseId"` // ID
    NurseCode         string                 `orm:"nurse_code" json:"nurseCode"`     // 编号
    NurseName         string                 `orm:"nurse_name" json:"nurseName"`     // 名称
    Sex               int                    `orm:"sex" json:"sex"`                  // 性别
    Birthday          *gtime.Time            `orm:"birthday" json:"birthday"`        // 出生年月
    AreaId            int                    `orm:"area_id" json:"areaId"`           // 院区
    DeptId            int64                  `orm:"dept_id" json:"deptId"`           // 科室
    LinkedDeptId      *LinkedNurseDepartment `orm:"with:dept_id=dept_id" json:"linkedDeptId"`
    StartDate         *gtime.Time            `orm:"start_date" json:"startDate"`     // 入职日期
    EducationId       int64                  `orm:"education_id" json:"educationId"` // 学历
    LinkedEducationId *LinkedNurseEducation  `orm:"with:education_id=education_id" json:"linkedEducationId"`
    TitleId           int64                  `orm:"title_id" json:"titleId"` // 职称
    LinkedTitleId     *LinkedNurseTitle      `orm:"with:title_id=title_id" json:"linkedTitleId"`
    StaffType         int                    `orm:"staff_type" json:"staffType"`      // 人员类别
    Note              string                 `orm:"note" json:"note"`                 // 备注
    EndDate           *gtime.Time            `orm:"end_date" json:"endDate"`          // 离职日期
    Status            int                    `orm:"status" json:"status"`             // 状态
    CreatedAt         *gtime.Time            `orm:"created_at" json:"createdAt"`      // 创建时间
    CreatedBy         string                 `orm:"created_by" json:"createdBy"`      // 创建用户
    UpdatedAt         *gtime.Time            `orm:"updated_at" json:"updatedAt"`      //
    UpdatedBy         string                 `orm:"updated_by" json:"updatedBy"`      // 修改用户
    DeletedAt         *gtime.Time            `orm:"deleted_at" json:"deletedAt"`      //
    DeletedBy         string                 `orm:"deleted_by" json:"deletedBy"`      // 删除用户
    SessionId         string                 `orm:"session_id" json:"sessionId"`      //
    CertEnddate       *gtime.Time            `orm:"cert_end_date" json:"certEnddate"` // 职业证书到期日期
    RegDate           *gtime.Time            `orm:"reg_date" json:"regDate"`          // 注册或变更日期
    CertNo            string                 `orm:"cert_no" json:"certNo"`            // 职业证书
    IdNo              string                 `orm:"id_no" json:"idNo"`                // 身份证号
    PositionId        int64                  `orm:"position_id" json:"positionId"`    // 职务
    LinkedPositionId  *LinkedNursePosition   `orm:"with:position_id=position_id" json:"linkedPositionId"`
}

type LinkedNursePosition struct {
    gmeta.Meta   `orm:"table:t_position"`
    PositionId   int64  `orm:"position_id" json:"positionId"`     //
    PositionName string `orm:"position_name" json:"positionName"` //
}

type LinkedNurseDepartment struct {
    gmeta.Meta `orm:"table:t_department"`
    DeptId     int64  `orm:"dept_id" json:"deptId"`     //
    DeptName   string `orm:"dept_name" json:"deptName"` //
}

type LinkedNurseEducation struct {
    gmeta.Meta    `orm:"table:t_education"`
    EducationId   int64  `orm:"education_id" json:"educationId"`     //
    EducationName string `orm:"education_name" json:"educationName"` //
}

type LinkedNurseTitle struct {
    gmeta.Meta `orm:"table:t_title"`
    TitleId    int64  `orm:"title_id" json:"titleId"`     //
    TitleName  string `orm:"title_name" json:"titleName"` //
}
