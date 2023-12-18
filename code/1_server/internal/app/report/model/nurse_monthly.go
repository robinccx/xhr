package model

type MonthlyTotalInfoRes struct {
	MonthStr  string `orm:"month_str" json:"monthStr"`   // 月份
	AreaId    int    `orm:"area_id" json:"areaId"`       // 院区
	StartQty  int    `orm:"start_qty" json:"startQty"`   // 月初人数
	LeaveQty  int    `orm:"leave_qty" json:"leaveQty"`   // 离职
	OutQty    int    `orm:"out_qty" json:"outQty"`       // 转出人数
	NewQty    int    `orm:"new_qty" json:"newQty"`       // 增减人数
	InQty     int    `orm:"in_qty" json:"inQty"`         // 转入人数
	AddQty    int    `orm:"add_qty" json:"addQty"`       // 增减人数
	EndQty    int    `orm:"end_qty" json:"endQty"`       // 月末人数
	LeaveList string `orm:"leave_list" json:"leaveList"` // 离职人员名单
	NewList   string `orm:"new_list" json:"newList"`     // 新入职人员名单
}

type MonthlyTotalChange struct {
	MonthStr      string `orm:"month_str" json:"monthStr"`            // 变更月份
	StartMonthStr string `orm:"start_month_str" json:"startMonthStr"` // 入院月份
	FromAreaId    int    `orm:"from_area_id" json:"fromAreaId"`       // 转出院区
	ToAreaId      int    `orm:"to_area_id" json:"toAreaId"`           // 转入院区
	Qty           int    `orm:"qty" json:"qty"`                       // 人数
}

type MonthlyEntityChange struct {
	MonthStr      string `orm:"month_str" json:"monthStr"`            // 变更月份
	StartMonthStr string `orm:"start_month_str" json:"startMonthStr"` // 入院月份
	EducationId   int64  `orm:"education_id" json:"education_id"`     // 学历
	TitleId       int64  `orm:"title_id" json:"title_id"`             // 职称
	FromEntityId  int64  `orm:"from_entity_id" json:"fromEntityId"`   // 原信息
	ToEntityId    int64  `orm:"to_entity_id" json:"toEntityId"`       // 新信息
	AreaId        int    `orm:"area_id" json:"areaId"`                // 转入院区
	FromAreaId    int    `orm:"from_area_id" json:"fromAreaId"`       // 转出院区
	ToAreaId      int    `orm:"to_area_id" json:"toAreaId"`           // 转入院区
	Qty           int64  `orm:"qty" json:"qty"`                       // 人数
}

type QuarterTotalChange struct {
	TransYear     int     `orm:"trans_year" json:"trans_year"`       // 变更年份
	QuarterIndex  int     `orm:"quarter_index" json:"quarterIndex"`  // 变更季度
	MonthStr      string  `orm:"month_str" json:"monthStr"`          // 变更月份
	StartMonthStr string  `orm:"start_month" json:"startMonth"`      // 入院月份
	StartYear     int     `orm:"start_year" json:"startYear"`        // 入院年份
	StartQuarter  int     `orm:"start_quarter" json:"start_quarter"` // 入院季度
	FromAreaId    int     `orm:"from_area_id" json:"fromAreaId"`     // 转出院区
	ToAreaId      int     `orm:"to_area_id" json:"toAreaId"`         // 转入院区
	NurseId       int64   `orm:"nurse_id" json:"nurseId"`            // ID
	DeptType      string  `orm:"dept_type" json:"deptType"`          // 科室类型
	TitleName     string  `orm:"title_name" json:"titleName"`        // 职称类别
	EducationSeq  int     `orm:"education_seq" json:"education_seq"` // 职称序号
	Y0            float64 `orm:"y0" json:"y0"`                       // 工作年数
	Y1            float64 `orm:"y1" json:"y1"`                       // 工作年数
	Y2            float64 `orm:"y2" json:"y2"`                       // 工作年数
	Y3            float64 `orm:"y3" json:"y3"`                       // 工作年数
}

type QuarterTotalInfoRes struct {
	Year      string `orm:"year" json:"year"`            // 年份
	ItemClass string `orm:"item_class" json:"itemClass"` // 分类名称
	ItemName  string `orm:"item_name" json:"itemName"`   // 分类名称
	ItemKey   string `orm:"item_key" json:"itemKey"`     // 分类名称

	StartQty1     int `orm:"start_qty1" json:"startQty1"`         // 年初杨浦
	StartQty2     int `orm:"start_qty2" json:"startQty2"`         // 年初安亭
	StartQtyTotal int `orm:"start_qtytotal" json:"startQtyTotal"` // 年初全院

	Quarter1Qty1     int `orm:"quarter1_qty1" json:"quarter1Qty1"`         // 第一季度杨浦
	Quarter1Qty2     int `orm:"quarter1_qty2" json:"quarter1Qty2"`         // 第一季度安亭
	Quarter1QtyTotal int `orm:"quarter1_qtytotal" json:"quarter1QtyTotal"` // 第一季度全院

	Quarter2Qty1     int `orm:"quarter2_qty1" json:"quarter2Qty1"`         // 第二季度杨浦
	Quarter2Qty2     int `orm:"quarter2_qty2" json:"quarter2Qty2"`         // 第二季度安亭
	Quarter2QtyTotal int `orm:"quarter2_qtytotal" json:"quarter2QtyTotal"` // 第二季度全院

	Quarter3Qty1     int `orm:"quarter3_qty1" json:"quarter3Qty1"`         // 第三季度杨浦
	Quarter3Qty2     int `orm:"quarter3_qty2" json:"quarter3Qty2"`         // 第三季度安亭
	Quarter3QtyTotal int `orm:"quarter3_qtytotal" json:"quarter3QtyTotal"` // 第三季度全院

	Quarter4Qty1     int `orm:"quarter4_qty1" json:"quarter4Qty1"`         // 第四季度杨浦
	Quarter4Qty2     int `orm:"quarter4_qty2" json:"quarter4Qty2"`         // 第四季度安亭
	Quarter4QtyTotal int `orm:"quarter4_qtytotal" json:"quarter4QtyTotal"` // 第四季度全院
}
