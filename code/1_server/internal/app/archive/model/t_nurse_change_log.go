// ==========================================================================
// GFast自动生成model操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/model/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================

package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

// NurseChangeLogInfoRes is the golang structure for table t_nurse_change_log.
type NurseChangeLogInfoRes struct {
	gmeta.Meta       `orm:"table:t_nurse_change_log"`
	Id               int64                           `orm:"id,primary" json:"id"`    //
	NurseId          int64                           `orm:"nurse_id" json:"nurseId"` //
	LinkedNurseId    *LinkedNurseChangeLogNurse      `orm:"with:nurse_id=nurse_id" json:"linkedNurseId"`
	ChangeType       int                             `orm:"change_type" json:"changeType"`  //
	FromAreaId       int                             `orm:"from_area_id" json:"fromAreaId"` //
	FromDeptId       int64                           `orm:"from_dept_id" json:"fromDeptId"` //
	LinkedFromDeptId *LinkedNurseChangeLogDepartment `orm:"with:dept_id=from_dept_id" json:"linkedFromDeptId"`
	FromEntityId     int64                           `orm:"from_entity_id" json:"fromEntityId"`     //
	FromEntityName   string                          `orm:"from_entity_name" json:"fromEntityName"` //
	ToAreaId         int                             `orm:"to_area_id" json:"toAreaId"`             //
	ToDeptId         int64                           `orm:"to_dept_id" json:"toDeptId"`             //
	LinkedToDeptId   *LinkedNurseChangeLogDepartment `orm:"with:dept_id=to_dept_id" json:"linkedToDeptId"`
	ToEntityId       int64                           `orm:"to_entity_id" json:"toEntityId"`     //
	ToEntityName     string                          `orm:"to_entity_name" json:"toEntityName"` //
	TransDate        *gtime.Time                     `orm:"trans_date" json:"transDate"`        //
	Status           int                             `orm:"status" json:"status"`               //
	CreatedAt        *gtime.Time                     `orm:"created_at" json:"createdAt"`        //
	CreatedBy        string                          `orm:"created_by" json:"createdBy"`        //
	UpdatedAt        *gtime.Time                     `orm:"updated_at" json:"updatedAt"`        //
	UpdatedBy        string                          `orm:"updated_by" json:"updatedBy"`        //
	DeletedAt        *gtime.Time                     `orm:"deleted_at" json:"deletedAt"`        //
	DeletedBy        string                          `orm:"deleted_by" json:"deletedBy"`        //
}

type LinkedNurseChangeLogDepartment struct {
	gmeta.Meta `orm:"table:t_department"`
	DeptId     int64 `orm:"dept_id" json:"deptId"` //
}

type LinkedNurseChangeLogNurse struct {
	gmeta.Meta `orm:"table:t_nurse"`
	NurseId    int64  `orm:"nurse_id" json:"nurseId"`     //
	NurseName  string `orm:"nurse_name" json:"nurseName"` //
}

type NurseChangeLogListRes struct {
	Id             int64                      `json:"id"`
	NurseId        int64                      `json:"nurseId"`
	LinkedNurseId  *LinkedNurseChangeLogNurse `orm:"with:nurse_id=nurse_id" json:"linkedNurseId"`
	ChangeType     int                        `json:"changeType"`
	FromEntityName string                     `json:"fromEntityName"`
	ToEntityName   string                     `json:"toEntityName"`
	TransDate      *gtime.Time                `json:"transDate"`
	CreatedAt      *gtime.Time                `json:"createdAt"`
	CreatedBy      string                     `json:"createdBy"`
}
