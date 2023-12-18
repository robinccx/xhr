package logic

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/report"
	"github.com/tiger1103/gfast/v3/internal/app/archive/dao"
	archiveEntity "github.com/tiger1103/gfast/v3/internal/app/archive/model/entity"
	commonDao "github.com/tiger1103/gfast/v3/internal/app/common/dao"
	commonEntity "github.com/tiger1103/gfast/v3/internal/app/common/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/report/model"
	"github.com/tiger1103/gfast/v3/internal/app/report/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"strconv"
	"strings"
	"time"
)

func init() {
	service.RegisterNurseMonthly(New())
}

func New() *sNurseMonthly {
	return &sNurseMonthly{}
}

type sNurseMonthly struct{}

func (s *sNurseMonthly) ListTotal(ctx context.Context, req *report.MonthlyTotalSearchReq) (res *report.MonthlyTotalSearchRes, err error) {
	res = new(report.MonthlyTotalSearchRes)

	err = g.Try(ctx, func(ctx context.Context) {
		var (
			monthIndex   int
			yearInt      int
			yearStr      string
			minMonth     string
			minDate      string
			maxDate      string
			startQty     int
			whereSql     string
			newDetail    []model.MonthlyTotalInfoRes
			LeaveDetail  []model.MonthlyTotalInfoRes
			ChangeDetail []model.MonthlyTotalChange
		)
		if req.RegStatus == "1" {
			whereSql = " and reg_date is not null "
		} else if req.RegStatus == "0" {
			whereSql = " and reg_date is null "
		}

		yearStr = req.Year
		if yearStr == "" {
			yearStr = strconv.Itoa(time.Now().Year())
			monthIndex = int(time.Now().Month())
		} else {
			if yearStr == strconv.Itoa(time.Now().Year()) {
				monthIndex = int(time.Now().Month())
			} else {
				monthIndex = 12
			}
		}
		yearInt, err = strconv.Atoi(yearStr)
		minMonth = yearStr + "-01"
		minDate = yearStr + "-01-01"
		maxDate = strconv.Itoa(yearInt+1) + "-01-01"

		if req.AreaId != "" {
			err = dao.Nurse.Ctx(ctx).
				Fields("area_id,DATE_FORMAT(start_date,'%Y-%m') as month_str,count(1)start_qty").
				Where("start_date>=? and start_date<? and area_id=?"+whereSql, minDate, maxDate, req.AreaId).
				Group("area_id,DATE_FORMAT(start_date,'%Y-%m')").Order("area_id,DATE_FORMAT(start_date,'%Y-%m')").Scan(&newDetail)
			liberr.ErrIsNil(ctx, err, "获取数据失败")

			err = dao.Nurse.Ctx(ctx).
				Fields("area_id,DATE_FORMAT(end_date,'%Y-%m') as month_str,count(1)start_qty,"+
					"GROUP_CONCAT(nurse_name order by nurse_name)leave_list").
				Where("status=2 and end_date>=? and end_date<? and area_id=?"+whereSql, minDate, maxDate, req.AreaId).
				Group("area_id,DATE_FORMAT(end_date,'%Y-%m')").Order("area_id,DATE_FORMAT(end_date,'%Y-%m')").Scan(&LeaveDetail)
			liberr.ErrIsNil(ctx, err, "获取数据失败")

		} else {

			err = dao.Nurse.Ctx(ctx).
				Fields("area_id,DATE_FORMAT(start_date,'%Y-%m') as month_str,count(1)start_qty").
				Where("start_date>=? and start_date<?"+whereSql, minDate, maxDate).
				Group("area_id,DATE_FORMAT(start_date,'%Y-%m')").Order("area_id,DATE_FORMAT(start_date,'%Y-%m')").Scan(&newDetail)
			liberr.ErrIsNil(ctx, err, "获取数据失败")

			err = dao.Nurse.Ctx(ctx).
				Fields("area_id,DATE_FORMAT(end_date,'%Y-%m') as month_str,count(1)start_qty,"+
					"GROUP_CONCAT(nurse_name order by nurse_name)leave_list").
				Where("status=2 and end_date>=? and end_date<?"+whereSql, minDate, maxDate).
				Group("area_id,DATE_FORMAT(end_date,'%Y-%m')").Order("area_id,DATE_FORMAT(end_date,'%Y-%m')").Scan(&LeaveDetail)
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}

		err = dao.Nurse.Ctx(ctx).InnerJoin("t_nurse_change_log a ", "t_nurse.nurse_id=a.nurse_id").
			Fields("DATE_FORMAT(a.trans_date,'%Y-%m') as month_str,"+
				"case when DATE_FORMAT(t_nurse.start_date,'%Y-%m')<'"+minMonth+"' then '' else DATE_FORMAT(t_nurse.start_date,'%Y-%m') end as start_month_str,"+
				"a.from_area_id,a.to_area_id ,count(1)qty").
			Where("a.change_type=1 and a.from_area_id<>a.to_area_id and a.trans_date>=?"+whereSql, minDate).
			Group("month_str,start_month_str,a.from_area_id,a.to_area_id").Order("month_str,start_month_str").Scan(&ChangeDetail)

		liberr.ErrIsNil(ctx, err, "获取数据失败")

		for areaIndex := 1; areaIndex <= 2; areaIndex++ {
			if req.AreaId != "" && req.AreaId != strconv.Itoa(areaIndex) {
				continue
			}
			startQty, err = dao.Nurse.Ctx(ctx).Count("start_date<? and area_id=? and (status=1 or status=2 and end_date>=?)",
				yearStr+"-01-01", areaIndex, yearStr+"-01-01")
			liberr.ErrIsNil(ctx, err, "获取期初数据失败")

			for i := 1; i <= monthIndex; i++ {
				resItem := &model.MonthlyTotalInfoRes{
					AreaId:   areaIndex,
					StartQty: startQty,
					LeaveQty: 0,
					NewQty:   0,
					AddQty:   0,
					EndQty:   0,
				}
				if i >= 10 {
					resItem.MonthStr = yearStr + "-" + strconv.Itoa(i)
				} else {
					resItem.MonthStr = yearStr + "-" + "0" + strconv.Itoa(i)
				}

				for _, detail := range newDetail {
					if detail.AreaId != areaIndex {
						continue
					}
					if resItem.MonthStr == detail.MonthStr {
						resItem.NewQty = detail.StartQty
						break
					} else if resItem.MonthStr < detail.MonthStr {
						break
					}
				}

				for _, detail := range LeaveDetail {
					if detail.AreaId != areaIndex {
						continue
					}
					if resItem.MonthStr == detail.MonthStr {
						resItem.LeaveQty = detail.StartQty
						resItem.LeaveList = detail.LeaveList
						break
					} else if resItem.MonthStr < detail.MonthStr {
						break
					}
				}
				resItem.AddQty = resItem.NewQty - resItem.LeaveQty
				resItem.EndQty = resItem.StartQty + resItem.AddQty
				startQty = startQty + resItem.AddQty
				res.List = append(res.List, resItem)
			}
		}

		for _, d := range ChangeDetail {
			for _, list := range res.List {
				if list.AreaId == d.FromAreaId {
					if d.MonthStr > list.MonthStr {
						if d.StartMonthStr == "" || d.StartMonthStr < list.MonthStr && d.StartMonthStr != "" {
							list.StartQty = list.StartQty + d.Qty
							list.EndQty = list.EndQty + d.Qty
						} else if d.StartMonthStr == list.MonthStr {
							list.NewQty = list.NewQty + d.Qty
							list.AddQty = list.AddQty + d.Qty
							list.EndQty = list.EndQty + d.Qty
						}

					} else if d.MonthStr == list.MonthStr {
						list.OutQty = list.OutQty + d.Qty
						list.AddQty = list.AddQty - d.Qty
						list.EndQty = list.EndQty - d.Qty
					}
				} else if list.AreaId == d.ToAreaId {
					if d.MonthStr > list.MonthStr {
						if d.StartMonthStr == "" || d.StartMonthStr < list.MonthStr && d.StartMonthStr != "" {
							list.StartQty = list.StartQty - d.Qty
							list.EndQty = list.EndQty - d.Qty
						} else if d.StartMonthStr == list.MonthStr {
							list.NewQty = list.NewQty - d.Qty
							list.AddQty = list.AddQty - d.Qty
							list.EndQty = list.EndQty - d.Qty
						}
					} else if d.MonthStr == list.MonthStr {
						list.InQty = list.InQty + d.Qty
						list.AddQty = list.AddQty + d.Qty
						list.EndQty = list.EndQty + d.Qty
					}
				}

			}
		}
	})
	return
}

func (s *sNurseMonthly) ListTitle(ctx context.Context, req *report.MonthlyTitleSearchReq) (res *report.MonthlyTitleSearchRes, err error) {
	res = new(report.MonthlyTitleSearchRes)

	err = g.Try(ctx, func(ctx context.Context) {
		var (
			monthIndex     int
			yearInt        int
			yearStr        string
			minMonth       string
			minDate        string
			maxDate        string
			whereSql       string
			fieldStr       string
			fieldLeaveStr  string
			fieldTotalStr  string
			titleList      []archiveEntity.Title
			nurseData      []map[string]interface{}
			nurseLeaveData []map[string]interface{}
			nurseStartData []map[string]interface{}
			ChangeDetail   []model.MonthlyEntityChange
			monthStr       string
			startIndex     int
			fieldName      string
			rst            gdb.Result
			qty            int64
		)
		if req.RegStatus == "1" {
			whereSql = " and reg_date is not null "
		} else if req.RegStatus == "0" {
			whereSql = " and reg_date is null "
		}

		yearStr = req.Year
		if yearStr == "" {
			yearStr = strconv.Itoa(time.Now().Year())
			monthIndex = int(time.Now().Month())
		} else {
			if yearStr == strconv.Itoa(time.Now().Year()) {
				monthIndex = int(time.Now().Month())
			} else {
				monthIndex = 12
			}
		}
		yearInt, err = strconv.Atoi(yearStr)
		minDate = yearStr + "-01-01"
		maxDate = strconv.Itoa(yearInt+1) + "-01-01"
		minMonth = yearStr + "-01"
		//获取职称列表，用于动态拼字段
		dao.Title.Ctx(ctx).Fields("title_id,title_name,seq").Order("seq").Scan(&titleList)
		liberr.ErrIsNil(ctx, err, "获取数据失败")

		// 获取期初数据和当前每月数据

		fieldTotalStr = "area_id,count(1)total_qty"

		fieldStr = "area_id,DATE_FORMAT(start_date,'%Y-%m') as month_str,count(1) total_qty"

		fieldLeaveStr = "area_id,DATE_FORMAT(end_date,'%Y-%m') as month_str,count(1)total_qty"
		for _, title := range titleList {
			fieldTotalStr = fieldTotalStr + fmt.Sprintf(",sum(case when title_id=%d then 1 else 0 end) as qty_%d",
				title.TitleId, title.TitleId)

			fieldStr = fieldStr + fmt.Sprintf(",sum(case when title_id=%d then 1 else 0 end) as qty_%d",
				title.TitleId, title.TitleId)

			fieldLeaveStr = fieldLeaveStr + fmt.Sprintf(",sum(case when title_id=%d then 1 else 0 end) as qty_%d",
				title.TitleId, title.TitleId)
		}

		// 获取期初数据
		if req.AreaId != "" {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldTotalStr).Where("start_date<? and area_id=? "+
				" and (status=1 or status=2 and end_date>=?)"+whereSql, minDate, req.AreaId, minDate).
				Group("area_id").Order("area_id").All()

			liberr.ErrIsNil(ctx, err, "获取数据失败")
		} else {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldTotalStr).Where("start_date<? "+
				" and (status=1 or status=2 and end_date>=?)"+whereSql, minDate, minDate).
				Group("area_id").Order("area_id").All()

			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}

		if rst != nil {
			for _, record := range rst {
				//detail := make(map[string]interface{})
				//for fieldName, fieldValue := range record {
				//	detail[fieldName] = fieldValue.Val()
				//}

				detail := map[string]interface{}{
					"area_id":   record["area_id"].Int(),
					"total_qty": record["total_qty"].Int64(),
				}

				for _, title := range titleList {
					fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
					detail[fieldName] = int64(record[fieldName].Float64())
				}
				nurseStartData = append(nurseStartData, detail)
			}
		}

		// 获取每月入职数据
		if req.AreaId != "" {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldStr).Where("start_date>=? and start_date<? and area_id=? "+whereSql, minDate, maxDate, req.AreaId).
				Group("area_id,DATE_FORMAT(start_date,'%Y-%m')").Order("area_id,DATE_FORMAT(start_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		} else {
			// 获取当前每月数据
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldStr).Where("start_date>=? and start_date<?"+whereSql, minDate, maxDate).
				Group("area_id,DATE_FORMAT(start_date,'%Y-%m')").Order("area_id,DATE_FORMAT(start_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}
		if rst != nil {
			for _, record := range rst {
				detail := map[string]interface{}{
					"area_id":   record["area_id"].Int(),
					"month_str": record["month_str"].String(),
					"total_qty": record["total_qty"].Int64(),
				}

				for _, title := range titleList {
					fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
					detail[fieldName] = int64(record[fieldName].Float64())
				}

				nurseData = append(nurseData, detail)
			}
		}

		// 获取每月离职数据
		if req.AreaId != "" {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldLeaveStr).Where(" status=2 and end_date>=? and end_date<? and area_id=? "+whereSql, minDate, maxDate, req.AreaId).
				Group("area_id,DATE_FORMAT(end_date,'%Y-%m')").Order("area_id,DATE_FORMAT(end_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		} else {
			// 获取当前每月数据
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldLeaveStr).Where(" status=2 and end_date>=? and end_date<?"+whereSql, minDate, maxDate).
				Group("area_id,DATE_FORMAT(end_date,'%Y-%m')").Order("area_id,DATE_FORMAT(end_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}
		if rst != nil {
			for _, record := range rst {
				detail := map[string]interface{}{
					"area_id":   record["area_id"].Int(),
					"month_str": record["month_str"].String(),
					"total_qty": record["total_qty"].Int64(),
				}

				for _, title := range titleList {
					fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
					detail[fieldName] = int64(record[fieldName].Float64())
				}
				nurseLeaveData = append(nurseLeaveData, detail)
			}
		}

		for areaIndex := 1; areaIndex <= 2; areaIndex++ {
			if req.AreaId != "" && req.AreaId != strconv.Itoa(areaIndex) {
				continue
			}
			startIndex = -1
			for index, detail := range nurseStartData {
				if detail["area_id"].(int) == areaIndex {
					startIndex = index
					break
				}
			}

			if startIndex < 0 {
				//如果没有期初数据，以零数量，补充
				detail := map[string]interface{}{
					"area_id":   areaIndex,
					"total_qty": int64(0),
				}

				for _, title := range titleList {
					fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
					detail[fieldName] = int64(0)
				}

				nurseStartData = append(nurseStartData, detail)

				startIndex = len(nurseStartData) - 1
			}

			for i := 1; i <= monthIndex; i++ {
				if i >= 10 {
					monthStr = yearStr + "-" + strconv.Itoa(i)
				} else {
					monthStr = yearStr + "-" + "0" + strconv.Itoa(i)
				}
				curDetail := map[string]interface{}{
					"month_str": monthStr,
					"area_id":   areaIndex,
					"total_qty": int64(0),
				}

				//总数量和分类数量，以期初记录数量赋值
				qty = nurseStartData[startIndex]["total_qty"].(int64)
				curDetail["total_qty"] = qty

				for _, title := range titleList {
					fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
					qty = nurseStartData[startIndex][fieldName].(int64)
					curDetail[fieldName] = qty
				}

				for _, detail := range nurseData {
					if detail["area_id"].(int) != areaIndex {
						continue
					}

					if detail["month_str"].(string) == monthStr {

						for _, title := range titleList {
							fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
							qty = detail[fieldName].(int64) + curDetail[fieldName].(int64)
							curDetail[fieldName] = qty
							nurseStartData[startIndex][fieldName] = qty
						}

						qty = detail["total_qty"].(int64) + curDetail["total_qty"].(int64)
						curDetail["total_qty"] = qty
						nurseStartData[startIndex]["total_qty"] = qty

					} else if detail["month_str"].(string) > monthStr {
						break
					}
				}

				for _, detail := range nurseLeaveData {
					if detail["area_id"].(int) != areaIndex {
						continue
					}

					if detail["month_str"].(string) == monthStr {

						for _, title := range titleList {
							fieldName = "qty_" + strconv.FormatInt(title.TitleId, 10)
							qty = curDetail[fieldName].(int64) - detail[fieldName].(int64)
							curDetail[fieldName] = qty
							nurseStartData[startIndex][fieldName] = qty
						}

						qty = curDetail["total_qty"].(int64) - detail["total_qty"].(int64)
						curDetail["total_qty"] = qty
						nurseStartData[startIndex]["total_qty"] = qty

					} else if detail["month_str"].(string) > monthStr {
						break
					}
				}
				res.List = append(res.List, curDetail)

			}
		}

		err = dao.Nurse.Ctx(ctx).InnerJoin("t_nurse_change_log a ", "t_nurse.nurse_id=a.nurse_id").
			Fields("t_nurse.area_id,a.from_area_id,a.to_area_id,DATE_FORMAT(a.trans_date,'%Y-%m') as month_str,"+
				"case when DATE_FORMAT(t_nurse.start_date,'%Y-%m')<'"+minMonth+"' then '' else DATE_FORMAT(t_nurse.start_date,'%Y-%m') end as start_month_str,"+
				"a.from_entity_id,a.to_entity_id,t_nurse.title_id,count(1)qty").
			Where("a.change_type=1 and a.from_area_id<>a.to_area_id and a.trans_date>=?"+whereSql, minDate).
			Group("t_nurse.area_id,t_nurse.education_id,month_str,start_month_str,a.from_entity_id,a.to_entity_id").
			Order("month_str,start_month_str").Scan(&ChangeDetail)

		liberr.ErrIsNil(ctx, err, "获取数据失败")

		for _, d := range ChangeDetail {
			for i, list := range res.List {
				if list["area_id"].(int) == d.FromAreaId {
					monthStr = list["month_str"].(string)
					if d.MonthStr >= monthStr {
						if d.StartMonthStr <= monthStr {
							fieldName = "qty_" + strconv.FormatInt(d.TitleId, 10)
							qty = list[fieldName].(int64) + d.Qty
							res.List[i][fieldName] = qty

							qty = res.List[i]["total_qty"].(int64) + d.Qty
							res.List[i]["total_qty"] = qty

						}
					}
				} else if list["area_id"].(int) == d.ToAreaId {
					monthStr = list["month_str"].(string)
					if d.MonthStr > monthStr {
						if d.StartMonthStr <= monthStr {
							fieldName = "qty_" + strconv.FormatInt(d.TitleId, 10)
							qty = list[fieldName].(int64) - d.Qty
							res.List[i][fieldName] = qty

							qty = res.List[i]["total_qty"].(int64) - d.Qty
							res.List[i]["total_qty"] = qty
						}
					}
				}
			}
		}
	})
	return
}

func (s *sNurseMonthly) ListEdu(ctx context.Context, req *report.MonthlyEduSearchReq) (res *report.MonthlyEduSearchRes, err error) {
	res = new(report.MonthlyEduSearchRes)

	err = g.Try(ctx, func(ctx context.Context) {
		var (
			monthIndex     int
			yearInt        int
			yearStr        string
			minMonth       string
			minDate        string
			maxDate        string
			whereSql       string
			fieldStr       string
			fieldLeaveStr  string
			fieldTotalStr  string
			eduList        []archiveEntity.Education
			nurseData      []map[string]interface{}
			nurseLeaveData []map[string]interface{}
			nurseStartData []map[string]interface{}
			ChangeDetail   []model.MonthlyEntityChange
			monthStr       string
			startIndex     int
			fieldName      string
			qty            int64
			rst            gdb.Result
		)
		if req.RegStatus == "1" {
			whereSql = " and reg_date is not null "
		} else if req.RegStatus == "0" {
			whereSql = " and reg_date is null "
		}

		yearStr = req.Year
		if yearStr == "" {
			yearStr = strconv.Itoa(time.Now().Year())
			monthIndex = int(time.Now().Month())
		} else {
			if yearStr == strconv.Itoa(time.Now().Year()) {
				monthIndex = int(time.Now().Month())
			} else {
				monthIndex = 12
			}
		}
		yearInt, err = strconv.Atoi(yearStr)
		minDate = yearStr + "-01-01"
		maxDate = strconv.Itoa(yearInt+1) + "-01-01"
		minMonth = yearStr + "-01"

		//获取职称列表，用于动态拼字段
		dao.Education.Ctx(ctx).Fields("education_id,education_name,seq").Order("seq").Scan(&eduList)
		liberr.ErrIsNil(ctx, err, "获取数据失败")

		// 获取期初数据和当前每月数据

		fieldTotalStr = "area_id,count(1)total_qty"

		fieldStr = "area_id,DATE_FORMAT(start_date,'%Y-%m') as month_str,count(1) total_qty"

		fieldLeaveStr = "area_id,DATE_FORMAT(end_date,'%Y-%m') as month_str,count(1)total_qty"
		for _, title := range eduList {
			fieldTotalStr = fieldTotalStr + fmt.Sprintf(",sum(case when education_id=%d then 1 else 0 end) as qty_%d",
				title.EducationId, title.EducationId)

			fieldStr = fieldStr + fmt.Sprintf(",sum(case when education_id=%d then 1 else 0 end) as qty_%d",
				title.EducationId, title.EducationId)

			fieldLeaveStr = fieldLeaveStr + fmt.Sprintf(",sum(case when education_id=%d then 1 else 0 end) as qty_%d",
				title.EducationId, title.EducationId)
		}

		// 获取期初数据
		if req.AreaId != "" {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldTotalStr).Where("start_date<? and area_id=? "+
				" and (status=1 or status=2 and end_date>=?)"+whereSql, minDate, req.AreaId, minDate).
				Group("area_id").Order("area_id").All()

			liberr.ErrIsNil(ctx, err, "获取数据失败")
		} else {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldTotalStr).Where("start_date<? "+
				" and (status=1 or status=2 and end_date>=?)"+whereSql, minDate, minDate).
				Group("area_id").Order("area_id").All()

			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}

		if rst != nil {
			for _, record := range rst {
				detail := map[string]interface{}{
					"area_id":   record["area_id"].Int(),
					"total_qty": record["total_qty"].Int64(),
				}

				for _, item := range eduList {
					fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
					detail[fieldName] = int64(record[fieldName].Float64())
				}
				nurseStartData = append(nurseStartData, detail)
			}
		}

		// 获取每月入职数据
		if req.AreaId != "" {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldStr).Where("start_date>=? and start_date<? and area_id=? "+whereSql, minDate, maxDate, req.AreaId).
				Group("area_id,DATE_FORMAT(start_date,'%Y-%m')").Order("area_id,DATE_FORMAT(start_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		} else {
			// 获取当前每月数据
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldStr).Where("start_date>=? and start_date<?"+whereSql, minDate, maxDate).
				Group("area_id,DATE_FORMAT(start_date,'%Y-%m')").Order("area_id,DATE_FORMAT(start_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}
		if rst != nil {
			for _, record := range rst {
				detail := map[string]interface{}{
					"area_id":   record["area_id"].Int(),
					"month_str": record["month_str"].String(),
					"total_qty": record["total_qty"].Int64(),
				}

				for _, item := range eduList {
					fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
					detail[fieldName] = int64(record[fieldName].Float64())
				}

				nurseData = append(nurseData, detail)
			}
		}

		// 获取每月离职数据
		if req.AreaId != "" {
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldLeaveStr).Where(" status=2 and end_date>=? and end_date<? and area_id=? "+whereSql, minDate, maxDate, req.AreaId).
				Group("area_id,DATE_FORMAT(end_date,'%Y-%m')").Order("area_id,DATE_FORMAT(end_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		} else {
			// 获取当前每月数据
			rst, err = dao.Nurse.Ctx(ctx).
				Fields(fieldLeaveStr).Where(" status=2 and end_date>=? and end_date<?"+whereSql, minDate, maxDate).
				Group("area_id,DATE_FORMAT(end_date,'%Y-%m')").Order("area_id,DATE_FORMAT(end_date,'%Y-%m')").All()
			liberr.ErrIsNil(ctx, err, "获取数据失败")
		}
		if rst != nil {
			for _, record := range rst {
				detail := map[string]interface{}{
					"area_id":   record["area_id"].Int(),
					"month_str": record["month_str"].String(),
					"total_qty": record["total_qty"].Int64(),
				}

				for _, item := range eduList {
					fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
					detail[fieldName] = int64(record[fieldName].Float64())
				}
				nurseLeaveData = append(nurseLeaveData, detail)
			}
		}

		for areaIndex := 1; areaIndex <= 2; areaIndex++ {
			if req.AreaId != "" && req.AreaId != strconv.Itoa(areaIndex) {
				continue
			}
			startIndex = -1
			for index, detail := range nurseStartData {
				if detail["area_id"].(int) == areaIndex {
					startIndex = index
					break
				}
			}

			if startIndex < 0 {
				//如果没有期初数据，以零数量，补充
				detail := map[string]interface{}{
					"area_id":   areaIndex,
					"total_qty": int64(0),
				}

				for _, item := range eduList {
					fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
					detail[fieldName] = int64(0)
				}

				nurseStartData = append(nurseStartData, detail)

				startIndex = len(nurseStartData) - 1
			}

			for i := 1; i <= monthIndex; i++ {
				if i >= 10 {
					monthStr = yearStr + "-" + strconv.Itoa(i)
				} else {
					monthStr = yearStr + "-" + "0" + strconv.Itoa(i)
				}
				curDetail := map[string]interface{}{
					"month_str": monthStr,
					"area_id":   areaIndex,
					"total_qty": int64(0),
				}

				//总数量和分类数量，以期初记录数量赋值
				qty = nurseStartData[startIndex]["total_qty"].(int64)
				curDetail["total_qty"] = qty

				for _, item := range eduList {
					fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
					qty = nurseStartData[startIndex][fieldName].(int64)
					curDetail[fieldName] = qty
				}

				for _, detail := range nurseData {
					if detail["area_id"].(int) != areaIndex {
						continue
					}

					if detail["month_str"].(string) == monthStr {

						for _, item := range eduList {
							fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
							qty = detail[fieldName].(int64) + curDetail[fieldName].(int64)
							curDetail[fieldName] = qty
							nurseStartData[startIndex][fieldName] = qty
						}

						qty = detail["total_qty"].(int64) + curDetail["total_qty"].(int64)
						curDetail["total_qty"] = qty
						nurseStartData[startIndex]["total_qty"] = qty

					} else if detail["month_str"].(string) > monthStr {
						break
					}
				}

				for _, detail := range nurseLeaveData {
					if detail["area_id"].(int) != areaIndex {
						continue
					}

					if detail["month_str"].(string) == monthStr {

						for _, item := range eduList {
							fieldName = "qty_" + strconv.FormatInt(item.EducationId, 10)
							qty = curDetail[fieldName].(int64) - detail[fieldName].(int64)
							curDetail[fieldName] = qty
							nurseStartData[startIndex][fieldName] = qty
						}

						qty = curDetail["total_qty"].(int64) - detail["total_qty"].(int64)
						curDetail["total_qty"] = qty
						nurseStartData[startIndex]["total_qty"] = qty

					} else if detail["month_str"].(string) > monthStr {
						break
					}
				}
				res.List = append(res.List, curDetail)

			}
		}

		err = dao.Nurse.Ctx(ctx).InnerJoin("t_nurse_change_log a ", "t_nurse.nurse_id=a.nurse_id").
			Fields("t_nurse.area_id,DATE_FORMAT(a.trans_date,'%Y-%m') as month_str,"+
				"case when DATE_FORMAT(t_nurse.start_date,'%Y-%m')<'"+minMonth+"' then '' else DATE_FORMAT(t_nurse.start_date,'%Y-%m') end as start_month_str,"+
				"a.from_entity_id,a.to_entity_id ,count(1)qty").
			Where("a.change_type=3 and a.from_entity_id<>a.to_entity_id and a.trans_date>=?"+whereSql, minDate).
			Group("t_nurse.area_id,month_str,start_month_str,a.from_entity_id,a.to_entity_id").Order("month_str,start_month_str").Scan(&ChangeDetail)

		liberr.ErrIsNil(ctx, err, "获取数据失败")

		for _, d := range ChangeDetail {
			for i, list := range res.List {
				if list["area_id"].(int) != d.AreaId {
					continue
				}
				monthStr = list["month_str"].(string)
				if d.MonthStr > monthStr {
					if d.StartMonthStr <= monthStr {
						fieldName = "qty_" + strconv.FormatInt(d.FromEntityId, 10)
						qty = list[fieldName].(int64) + d.Qty
						res.List[i][fieldName] = qty

						fieldName = "qty_" + strconv.FormatInt(d.ToEntityId, 10)
						qty = list[fieldName].(int64) - d.Qty
						res.List[i][fieldName] = qty
					}
				} else if d.MonthStr == monthStr {
					fieldName = "qty_" + strconv.FormatInt(d.FromEntityId, 10)
					qty = list[fieldName].(int64) + d.Qty
					res.List[i][fieldName] = qty
				}
			}
		}

		err = dao.Nurse.Ctx(ctx).InnerJoin("t_nurse_change_log a ", "t_nurse.nurse_id=a.nurse_id").
			Fields("t_nurse.area_id,a.from_area_id,a.to_area_id,DATE_FORMAT(a.trans_date,'%Y-%m') as month_str,"+
				"case when DATE_FORMAT(t_nurse.start_date,'%Y-%m')<'"+minMonth+"' then '' else DATE_FORMAT(t_nurse.start_date,'%Y-%m') end as start_month_str,"+
				"a.from_entity_id,a.to_entity_id,t_nurse.education_id,count(1)qty").
			Where("a.change_type=1 and a.from_area_id<>a.to_area_id and a.trans_date>=?"+whereSql, minDate).
			Group("t_nurse.area_id,t_nurse.education_id,month_str,start_month_str,a.from_entity_id,a.to_entity_id").
			Order("month_str,start_month_str").Scan(&ChangeDetail)

		liberr.ErrIsNil(ctx, err, "获取数据失败")

		for _, d := range ChangeDetail {
			for i, list := range res.List {
				if list["area_id"].(int) == d.FromAreaId {
					monthStr = list["month_str"].(string)
					if d.MonthStr >= monthStr {
						if d.StartMonthStr <= monthStr {
							fieldName = "qty_" + strconv.FormatInt(d.EducationId, 10)
							qty = list[fieldName].(int64) + d.Qty
							res.List[i][fieldName] = qty

							qty = res.List[i]["total_qty"].(int64) + d.Qty
							res.List[i]["total_qty"] = qty

						}
					}
				} else if list["area_id"].(int) == d.ToAreaId {
					monthStr = list["month_str"].(string)
					if d.MonthStr > monthStr {
						if d.StartMonthStr <= monthStr {
							fieldName = "qty_" + strconv.FormatInt(d.EducationId, 10)
							qty = list[fieldName].(int64) - d.Qty
							res.List[i][fieldName] = qty

							qty = res.List[i]["total_qty"].(int64) - d.Qty
							res.List[i]["total_qty"] = qty
						}
					}
				}
			}
		}
	})
	return
}

func (s *sNurseMonthly) getQuarterTotalAges(ctx context.Context, year string) (rst []*model.QuarterTotalInfoRes, err error) {
	var (
		fieldSql  string
		whereSql  string
		itemClass string
		curYear   string
		endDate   string
		lastYear  int
		nextYear  int
		curMonth  int
		detail    []model.QuarterTotalInfoRes
	)

	itemClass = "人力资源结构--工作年限相关数据"
	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季初<1年资护士人数",
		ItemKey:   "411",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季末<1年资护士人数",
		ItemKey:   "412",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季初1≤y<2年资护士人数",
		ItemKey:   "421",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季末1≤y<2年资护士人数",
		ItemKey:   "422",
	})
	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季初2≤y<5年资护士人数",
		ItemKey:   "431",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季末2≤y<5年资护士人数",
		ItemKey:   "432",
	})
	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季初5≤y<10年资护士人数",
		ItemKey:   "441",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季末5≤y<10年资护士人数",
		ItemKey:   "442",
	})
	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季初10≤y<20年资护士人数",
		ItemKey:   "451",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季末10≤y<20年资护士人数",
		ItemKey:   "452",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季初≥20年资护士人数",
		ItemKey:   "461",
	})

	rst = append(rst, &model.QuarterTotalInfoRes{
		Year:      year,
		ItemClass: itemClass,
		ItemName:  "季末≥20年资护士人数",
		ItemKey:   "462",
	})

	lastYear, err = strconv.Atoi(year)
	nextYear = lastYear + 1
	lastYear = lastYear - 1

	curYear = strconv.Itoa(time.Now().Year())
	curMonth, err = strconv.Atoi(time.Now().Format("01"))

	//一季度 期初
	endDate = strconv.Itoa(lastYear) + "-12-31"
	fieldSql = fmt.Sprintf("case when DATEDIFF('%s',start_date)/365.0<1 then 1 "+
		" when DATEDIFF('%s',start_date)/365.0<2 then 2 "+
		" when DATEDIFF('%s',start_date)/365.0<5 then 3 "+
		" when DATEDIFF('%s',start_date)/365.0<10 then 4 "+
		" when DATEDIFF('%s',start_date)/365.0<20 then 5 ", endDate, endDate, endDate, endDate, endDate)
	fieldSql = fieldSql + " else 6 end as item_key,"
	fieldSql = fieldSql + "sum(case when area_id=1 then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 then 1  else 0 end )quarter1_qty2"
	whereSql = fmt.Sprintf("reg_date is not null and start_date<'%s' and (status=1 or status=2 and end_date>='%s')", year+"-01-01", year+"-01-01")
	err = dao.Nurse.Ctx(ctx).Fields(fieldSql).
		Where(whereSql).Group("item_key").Scan(&detail)
	liberr.ErrIsNilEx(ctx, err, "获取数据错误")

	for _, d := range detail {
		switch d.ItemKey {
		case "1":
			rst[0].Quarter1Qty1 = d.Quarter1Qty1
			rst[0].Quarter1Qty2 = d.Quarter1Qty2
		case "2":
			rst[2].Quarter1Qty1 = d.Quarter1Qty1
			rst[2].Quarter1Qty2 = d.Quarter1Qty2
		case "3":
			rst[4].Quarter1Qty1 = d.Quarter1Qty1
			rst[4].Quarter1Qty2 = d.Quarter1Qty2
		case "4":
			rst[6].Quarter1Qty1 = d.Quarter1Qty1
			rst[6].Quarter1Qty2 = d.Quarter1Qty2
		case "5":
			rst[8].Quarter1Qty1 = d.Quarter1Qty1
			rst[8].Quarter1Qty2 = d.Quarter1Qty2
		case "6":
			rst[10].Quarter1Qty1 = d.Quarter1Qty1
			rst[10].Quarter1Qty2 = d.Quarter1Qty2
		}
	}

	//一季度 期末
	endDate = year + "-03-31"
	fieldSql = fmt.Sprintf("case when DATEDIFF('%s',start_date)/365.0<1 then 1 "+
		" when DATEDIFF('%s',start_date)/365.0<2 then 2 "+
		" when DATEDIFF('%s',start_date)/365.0<5 then 3 "+
		" when DATEDIFF('%s',start_date)/365.0<10 then 4 "+
		" when DATEDIFF('%s',start_date)/365.0<20 then 5 ", endDate, endDate, endDate, endDate, endDate)
	fieldSql = fieldSql + " else 6 end as item_key,"
	fieldSql = fieldSql + "sum(case when area_id=1 then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 then 1  else 0 end )quarter1_qty2"
	whereSql = fmt.Sprintf("reg_date is not null and start_date<'%s' and (status=1 or status=2 and end_date>='%s')", year+"-04-01", year+"-04-01")
	err = dao.Nurse.Ctx(ctx).Fields(fieldSql).
		Where(whereSql).Group("item_key").Scan(&detail)
	liberr.ErrIsNilEx(ctx, err, "获取数据错误")

	for _, d := range detail {
		switch d.ItemKey {
		case "1":
			rst[1].Quarter1Qty1 = d.Quarter1Qty1
			rst[1].Quarter1Qty2 = d.Quarter1Qty2

			rst[0].Quarter2Qty1 = d.Quarter1Qty1
			rst[0].Quarter2Qty2 = d.Quarter1Qty2
		case "2":
			rst[3].Quarter1Qty1 = d.Quarter1Qty1
			rst[3].Quarter1Qty2 = d.Quarter1Qty2

			rst[2].Quarter2Qty1 = d.Quarter1Qty1
			rst[2].Quarter2Qty2 = d.Quarter1Qty2
		case "3":
			rst[5].Quarter1Qty1 = d.Quarter1Qty1
			rst[5].Quarter1Qty2 = d.Quarter1Qty2

			rst[4].Quarter2Qty1 = d.Quarter1Qty1
			rst[4].Quarter2Qty2 = d.Quarter1Qty2
		case "4":
			rst[7].Quarter1Qty1 = d.Quarter1Qty1
			rst[7].Quarter1Qty2 = d.Quarter1Qty2

			rst[6].Quarter2Qty1 = d.Quarter1Qty1
			rst[6].Quarter2Qty2 = d.Quarter1Qty2
		case "5":
			rst[9].Quarter1Qty1 = d.Quarter1Qty1
			rst[9].Quarter1Qty2 = d.Quarter1Qty2

			rst[8].Quarter2Qty1 = d.Quarter1Qty1
			rst[8].Quarter2Qty2 = d.Quarter1Qty2
		case "6":
			rst[11].Quarter1Qty1 = d.Quarter1Qty1
			rst[11].Quarter1Qty2 = d.Quarter1Qty2

			rst[10].Quarter2Qty1 = d.Quarter1Qty1
			rst[10].Quarter2Qty2 = d.Quarter1Qty2
		}
	}

	if year == curYear && curMonth <= 3 {
		return
	}
	//二季度 期末
	endDate = year + "-06-30"
	fieldSql = fmt.Sprintf("case when DATEDIFF('%s',start_date)/365.0<1 then 1 "+
		" when DATEDIFF('%s',start_date)/365.0<2 then 2 "+
		" when DATEDIFF('%s',start_date)/365.0<5 then 3 "+
		" when DATEDIFF('%s',start_date)/365.0<10 then 4 "+
		" when DATEDIFF('%s',start_date)/365.0<20 then 5 ", endDate, endDate, endDate, endDate, endDate)
	fieldSql = fieldSql + " else 6 end as item_key,"
	fieldSql = fieldSql + "sum(case when area_id=1 then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 then 1  else 0 end )quarter1_qty2"
	whereSql = fmt.Sprintf("reg_date is not null and start_date<'%s' and (status=1 or status=2 and end_date>='%s')", year+"-07-01", year+"-07-01")
	err = dao.Nurse.Ctx(ctx).Fields(fieldSql).
		Where(whereSql).Group("item_key").Scan(&detail)
	liberr.ErrIsNilEx(ctx, err, "获取数据错误")

	for _, d := range detail {
		switch d.ItemKey {
		case "1":
			rst[1].Quarter2Qty1 = d.Quarter1Qty1
			rst[1].Quarter2Qty2 = d.Quarter1Qty2

			rst[0].Quarter3Qty1 = d.Quarter1Qty1
			rst[0].Quarter3Qty2 = d.Quarter1Qty2
		case "2":
			rst[3].Quarter2Qty1 = d.Quarter1Qty1
			rst[3].Quarter2Qty2 = d.Quarter1Qty2

			rst[2].Quarter3Qty1 = d.Quarter1Qty1
			rst[2].Quarter3Qty2 = d.Quarter1Qty2
		case "3":
			rst[5].Quarter2Qty1 = d.Quarter1Qty1
			rst[5].Quarter2Qty2 = d.Quarter1Qty2

			rst[4].Quarter3Qty1 = d.Quarter1Qty1
			rst[4].Quarter3Qty2 = d.Quarter1Qty2
		case "4":
			rst[7].Quarter2Qty1 = d.Quarter1Qty1
			rst[7].Quarter2Qty2 = d.Quarter1Qty2

			rst[6].Quarter3Qty1 = d.Quarter1Qty1
			rst[6].Quarter3Qty2 = d.Quarter1Qty2
		case "5":
			rst[9].Quarter2Qty1 = d.Quarter1Qty1
			rst[9].Quarter2Qty2 = d.Quarter1Qty2

			rst[8].Quarter3Qty1 = d.Quarter1Qty1
			rst[8].Quarter3Qty2 = d.Quarter1Qty2
		case "6":
			rst[11].Quarter2Qty1 = d.Quarter1Qty1
			rst[11].Quarter2Qty2 = d.Quarter1Qty2

			rst[10].Quarter3Qty1 = d.Quarter1Qty1
			rst[10].Quarter3Qty2 = d.Quarter1Qty2
		}
	}

	if year == curYear && curMonth <= 6 {
		return
	}

	//三季度 期末
	endDate = year + "-09-31"
	fieldSql = fmt.Sprintf("case when DATEDIFF('%s',start_date)/365.0<1 then 1 "+
		" when DATEDIFF('%s',start_date)/365.0<2 then 2 "+
		" when DATEDIFF('%s',start_date)/365.0<5 then 3 "+
		" when DATEDIFF('%s',start_date)/365.0<10 then 4 "+
		" when DATEDIFF('%s',start_date)/365.0<20 then 5 ", endDate, endDate, endDate, endDate, endDate)
	fieldSql = fieldSql + " else 6 end as item_key,"
	fieldSql = fieldSql + "sum(case when area_id=1 then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 then 1  else 0 end )quarter1_qty2"
	whereSql = fmt.Sprintf("reg_date is not null and start_date<'%s' and (status=1 or status=2 and end_date>='%s')", year+"-10-01", year+"-10-01")
	err = dao.Nurse.Ctx(ctx).Fields(fieldSql).
		Where(whereSql).Group("item_key").Scan(&detail)
	liberr.ErrIsNilEx(ctx, err, "获取数据错误")

	for _, d := range detail {
		switch d.ItemKey {
		case "1":
			rst[1].Quarter3Qty1 = d.Quarter1Qty1
			rst[1].Quarter3Qty2 = d.Quarter1Qty2

			rst[0].Quarter4Qty1 = d.Quarter1Qty1
			rst[0].Quarter4Qty2 = d.Quarter1Qty2
		case "2":
			rst[3].Quarter3Qty1 = d.Quarter1Qty1
			rst[3].Quarter3Qty2 = d.Quarter1Qty2

			rst[2].Quarter4Qty1 = d.Quarter1Qty1
			rst[2].Quarter4Qty2 = d.Quarter1Qty2
		case "3":
			rst[5].Quarter3Qty1 = d.Quarter1Qty1
			rst[5].Quarter3Qty2 = d.Quarter1Qty2

			rst[4].Quarter4Qty1 = d.Quarter1Qty1
			rst[4].Quarter4Qty2 = d.Quarter1Qty2
		case "4":
			rst[7].Quarter3Qty1 = d.Quarter1Qty1
			rst[7].Quarter3Qty2 = d.Quarter1Qty2

			rst[6].Quarter4Qty1 = d.Quarter1Qty1
			rst[6].Quarter4Qty2 = d.Quarter1Qty2
		case "5":
			rst[9].Quarter3Qty1 = d.Quarter1Qty1
			rst[9].Quarter3Qty2 = d.Quarter1Qty2

			rst[8].Quarter4Qty1 = d.Quarter1Qty1
			rst[8].Quarter4Qty2 = d.Quarter1Qty2
		case "6":
			rst[11].Quarter3Qty1 = d.Quarter1Qty1
			rst[11].Quarter3Qty2 = d.Quarter1Qty2

			rst[10].Quarter4Qty1 = d.Quarter1Qty1
			rst[10].Quarter4Qty2 = d.Quarter1Qty2
		}
	}

	if year == curYear && curMonth <= 6 {
		return
	}

	//四季度 期末
	endDate = year + "-12-31"
	fieldSql = fmt.Sprintf("case when DATEDIFF('%s',start_date)/365.0<1 then 1 "+
		" when DATEDIFF('%s',start_date)/365.0<2 then 2 "+
		" when DATEDIFF('%s',start_date)/365.0<5 then 3 "+
		" when DATEDIFF('%s',start_date)/365.0<10 then 4 "+
		" when DATEDIFF('%s',start_date)/365.0<20 then 5 ", endDate, endDate, endDate, endDate, endDate)
	fieldSql = fieldSql + " else 6 end as item_key,"
	fieldSql = fieldSql + "sum(case when area_id=1 then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 then 1  else 0 end )quarter1_qty2"
	whereSql = fmt.Sprintf("reg_date is not null and start_date<'%s' and (status=1 or status=2 and end_date>='%s')",
		strconv.Itoa(nextYear)+"-01-01", strconv.Itoa(nextYear)+"-01-01")
	err = dao.Nurse.Ctx(ctx).Fields(fieldSql).
		Where(whereSql).Group("item_key").Scan(&detail)
	liberr.ErrIsNilEx(ctx, err, "获取数据错误")

	for _, d := range detail {
		switch d.ItemKey {
		case "1":
			rst[1].Quarter4Qty1 = d.Quarter1Qty1
			rst[1].Quarter4Qty2 = d.Quarter1Qty2
		case "2":
			rst[3].Quarter4Qty1 = d.Quarter1Qty1
			rst[3].Quarter4Qty2 = d.Quarter1Qty2
		case "3":
			rst[5].Quarter4Qty1 = d.Quarter1Qty1
			rst[5].Quarter4Qty2 = d.Quarter1Qty2

		case "4":
			rst[7].Quarter4Qty1 = d.Quarter1Qty1
			rst[7].Quarter4Qty2 = d.Quarter1Qty2
		case "5":
			rst[9].Quarter4Qty1 = d.Quarter1Qty1
			rst[9].Quarter4Qty2 = d.Quarter1Qty2
		case "6":
			rst[11].Quarter4Qty1 = d.Quarter1Qty1
			rst[11].Quarter4Qty2 = d.Quarter1Qty2
		}
	}

	return
}
func (s *sNurseMonthly) getQuarterTotal(ctx context.Context, itemClass, itemName, itemKey, year, fieldSql, groupSql, sortSql,
	whereSql string) (rst []*model.QuarterTotalInfoRes, err error) {

	var (
		detail          []model.QuarterTotalInfoRes
		nextYear        int
		quarter1MinDate string
		quarter1MaxDate string
		quarter2MaxDate string
		quarter3MaxDate string
		quarter4MaxDate string
		fieldStr        string
		whereStr        string
	)
	nextYear, err = strconv.Atoi(year)
	nextYear = nextYear + 1

	quarter1MinDate = year + "-01-01"

	quarter1MaxDate = year + "-04-01"

	quarter2MaxDate = year + "-07-01"

	quarter3MaxDate = year + "-10-01"
	quarter4MaxDate = strconv.Itoa(nextYear) + "-01-01"

	fieldStr = fieldSql + "sum(case when area_id=1 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )start_qty1," +
		"sum(case when area_id=2 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )start_qty2," +
		"sum(case when start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )start_qtytotal," +
		"sum(case when area_id=1 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter1_qty2," +
		"sum(case when start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter1_qtytotal," +
		"sum(case when area_id=1 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter2_qty1," +
		"sum(case when area_id=2 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter2_qty2," +
		"sum(case when start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter2_qtytotal," +
		"sum(case when area_id=1 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter3_qty1," +
		"sum(case when area_id=2 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter3_qty2," +
		"sum(case when start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter3_qtytotal," +
		"sum(case when area_id=1 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter4_qty1," +
		"sum(case when area_id=2 and start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter4_qty2," +
		"sum(case when start_date<'%s' and (status=1 or status=2 and end_date>='%s') then 1  else 0 end )quarter4_qtytotal"
	fieldStr = fmt.Sprintf(fieldStr,
		quarter1MinDate, quarter1MinDate,
		quarter1MinDate, quarter1MinDate,
		quarter1MinDate, quarter1MinDate,
		quarter1MaxDate, quarter1MaxDate,
		quarter1MaxDate, quarter1MaxDate,
		quarter1MaxDate, quarter1MaxDate,
		quarter2MaxDate, quarter2MaxDate,
		quarter2MaxDate, quarter2MaxDate,
		quarter2MaxDate, quarter2MaxDate,
		quarter3MaxDate, quarter3MaxDate,
		quarter3MaxDate, quarter3MaxDate,
		quarter3MaxDate, quarter3MaxDate,
		quarter4MaxDate, quarter4MaxDate,
		quarter4MaxDate, quarter4MaxDate,
		quarter4MaxDate, quarter4MaxDate)
	whereStr = " reg_date is not null and start_date<?" + whereSql
	if groupSql == "" {
		err = dao.Nurse.Ctx(ctx).Fields(fieldStr).
			Where(whereStr, quarter4MaxDate).Scan(&detail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")

		for _, d := range detail {
			//rst []*model.QuarterTotalInfoRes
			item := &model.QuarterTotalInfoRes{
				ItemClass: itemClass,
				ItemName:  "季初" + itemName,
				ItemKey:   itemKey + "1",

				Quarter1Qty1:     d.StartQty1,
				Quarter1Qty2:     d.StartQty2,
				Quarter1QtyTotal: d.StartQtyTotal,

				Quarter2Qty1:     d.Quarter1Qty1,
				Quarter2Qty2:     d.Quarter1Qty2,
				Quarter2QtyTotal: d.Quarter1QtyTotal,

				Quarter3Qty1:     d.Quarter2Qty1,
				Quarter3Qty2:     d.Quarter2Qty2,
				Quarter3QtyTotal: d.Quarter2QtyTotal,

				Quarter4Qty1:     d.Quarter3Qty1,
				Quarter4Qty2:     d.Quarter3Qty2,
				Quarter4QtyTotal: d.Quarter3QtyTotal,
			}
			rst = append(rst, item)

			item = &model.QuarterTotalInfoRes{
				ItemClass:        itemClass,
				ItemName:         "季末" + itemName,
				ItemKey:          itemKey + "2",
				Quarter1Qty1:     d.Quarter1Qty1,
				Quarter1Qty2:     d.Quarter1Qty2,
				Quarter1QtyTotal: d.Quarter1QtyTotal,

				Quarter2Qty1:     d.Quarter2Qty1,
				Quarter2Qty2:     d.Quarter2Qty2,
				Quarter2QtyTotal: d.Quarter2QtyTotal,

				Quarter3Qty1:     d.Quarter3Qty1,
				Quarter3Qty2:     d.Quarter3Qty2,
				Quarter3QtyTotal: d.Quarter3QtyTotal,

				Quarter4Qty1:     d.Quarter4Qty1,
				Quarter4Qty2:     d.Quarter4Qty2,
				Quarter4QtyTotal: d.Quarter4QtyTotal,
			}
			rst = append(rst, item)
		}
	} else {

		err = dao.Nurse.Ctx(ctx).Fields(fieldStr).
			Where(whereStr, quarter4MaxDate).Group(groupSql).Order(sortSql).Scan(&detail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")

		for _, d := range detail {
			//rst []*model.QuarterTotalInfoRes
			item := &model.QuarterTotalInfoRes{
				ItemClass:        itemClass,
				ItemName:         "季初" + d.ItemName,
				ItemKey:          itemKey + d.ItemKey + "1",
				Quarter1Qty1:     d.StartQty1,
				Quarter1Qty2:     d.StartQty2,
				Quarter1QtyTotal: d.StartQtyTotal,

				Quarter2Qty1:     d.Quarter1Qty1,
				Quarter2Qty2:     d.Quarter1Qty2,
				Quarter2QtyTotal: d.Quarter1QtyTotal,

				Quarter3Qty1:     d.Quarter2Qty1,
				Quarter3Qty2:     d.Quarter2Qty2,
				Quarter3QtyTotal: d.Quarter2QtyTotal,

				Quarter4Qty1:     d.Quarter3Qty1,
				Quarter4Qty2:     d.Quarter3Qty2,
				Quarter4QtyTotal: d.Quarter3QtyTotal,
			}
			rst = append(rst, item)

			item = &model.QuarterTotalInfoRes{
				ItemClass:        itemClass,
				ItemName:         "季末" + d.ItemName,
				ItemKey:          itemKey + d.ItemKey + "2",
				Quarter1Qty1:     d.Quarter1Qty1,
				Quarter1Qty2:     d.Quarter1Qty2,
				Quarter1QtyTotal: d.Quarter1QtyTotal,

				Quarter2Qty1:     d.Quarter2Qty1,
				Quarter2Qty2:     d.Quarter2Qty2,
				Quarter2QtyTotal: d.Quarter2QtyTotal,

				Quarter3Qty1:     d.Quarter3Qty1,
				Quarter3Qty2:     d.Quarter3Qty2,
				Quarter3QtyTotal: d.Quarter3QtyTotal,

				Quarter4Qty1:     d.Quarter4Qty1,
				Quarter4Qty2:     d.Quarter4Qty2,
				Quarter4QtyTotal: d.Quarter4QtyTotal,
			}
			rst = append(rst, item)
		}
	}

	return
}

func (s *sNurseMonthly) getQuarterLevel(ctx context.Context, itemClass, itemName, year, fieldSql, groupSql, sortSql,
	whereSql string) (rst []*model.QuarterTotalInfoRes, err error) {

	var (
		detail          []model.QuarterTotalInfoRes
		nextYear        int
		quarter1MinDate string
		quarter1MaxDate string
		quarter2MaxDate string
		quarter3MaxDate string
		quarter4MaxDate string
		fieldStr        string
		whereStr        string
	)
	nextYear, err = strconv.Atoi(year)
	nextYear = nextYear + 1

	quarter1MinDate = year + "-01-01"

	quarter1MaxDate = year + "-04-01"

	quarter2MaxDate = year + "-07-01"

	quarter3MaxDate = year + "-10-01"
	quarter4MaxDate = strconv.Itoa(nextYear) + "-01-01"

	fieldStr = fieldSql + "sum(case when area_id=1 and end_date>='%s' and end_date<'%s' then 1  else 0 end )quarter1_qty1," +
		"sum(case when area_id=2 and end_date>='%s' and end_date<'%s' then 1  else 0 end )quarter1_qty2," +
		"sum(case when end_date>='%s'  and end_date<'%s' then 1  else 0 end )quarter1_qtytotal," +
		"sum(case when area_id=1 and end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter2_qty1," +
		"sum(case when area_id=2 and end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter2_qty2," +
		"sum(case when end_date>='%s' and end_date<'%s' then 1  else 0 end )quarter2_qtytotal," +
		"sum(case when area_id=1 and end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter3_qty1," +
		"sum(case when area_id=2 and end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter3_qty2," +
		"sum(case when end_date>='%s' and end_date<'%s' then 1  else 0 end )quarter3_qtytotal," +
		"sum(case when area_id=1 and end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter4_qty1," +
		"sum(case when area_id=2 and end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter4_qty2," +
		"sum(case when end_date>='%s' and end_date<'%s'  then 1  else 0 end )quarter4_qtytotal"
	fieldStr = fmt.Sprintf(fieldStr,
		quarter1MinDate, quarter1MaxDate,
		quarter1MinDate, quarter1MaxDate,
		quarter1MinDate, quarter1MaxDate,

		quarter1MaxDate, quarter2MaxDate,
		quarter1MaxDate, quarter2MaxDate,
		quarter1MaxDate, quarter2MaxDate,

		quarter2MaxDate, quarter3MaxDate,
		quarter2MaxDate, quarter3MaxDate,
		quarter2MaxDate, quarter3MaxDate,

		quarter3MaxDate, quarter4MaxDate,
		quarter3MaxDate, quarter4MaxDate,
		quarter3MaxDate, quarter4MaxDate)
	whereStr = " reg_date is not null and status=2 and end_date>=? and end_date<?" + whereSql
	if groupSql == "" {
		err = dao.Nurse.Ctx(ctx).Fields(fieldStr).
			Where(whereStr, quarter1MinDate, quarter4MaxDate).Scan(&detail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")

		for _, d := range detail {

			item := &model.QuarterTotalInfoRes{
				ItemClass:        itemClass,
				ItemName:         itemName,
				Quarter1Qty1:     d.Quarter1Qty1,
				Quarter1Qty2:     d.Quarter1Qty2,
				Quarter1QtyTotal: d.Quarter1QtyTotal,

				Quarter2Qty1:     d.Quarter2Qty1,
				Quarter2Qty2:     d.Quarter2Qty2,
				Quarter2QtyTotal: d.Quarter2QtyTotal,

				Quarter3Qty1:     d.Quarter3Qty1,
				Quarter3Qty2:     d.Quarter3Qty2,
				Quarter3QtyTotal: d.Quarter3QtyTotal,

				Quarter4Qty1:     d.Quarter4Qty1,
				Quarter4Qty2:     d.Quarter4Qty2,
				Quarter4QtyTotal: d.Quarter4QtyTotal,
			}
			rst = append(rst, item)
		}
	} else {

		err = dao.Nurse.Ctx(ctx).Fields(fieldStr).
			Where(whereStr, quarter1MinDate, quarter4MaxDate).Group(groupSql).Order(sortSql).Scan(&detail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")

		for _, d := range detail {

			item := &model.QuarterTotalInfoRes{
				ItemClass:        itemClass,
				ItemName:         d.ItemName,
				Quarter1Qty1:     d.Quarter1Qty1,
				Quarter1Qty2:     d.Quarter1Qty2,
				Quarter1QtyTotal: d.Quarter1QtyTotal,

				Quarter2Qty1:     d.Quarter2Qty1,
				Quarter2Qty2:     d.Quarter2Qty2,
				Quarter2QtyTotal: d.Quarter2QtyTotal,

				Quarter3Qty1:     d.Quarter3Qty1,
				Quarter3Qty2:     d.Quarter3Qty2,
				Quarter3QtyTotal: d.Quarter3QtyTotal,

				Quarter4Qty1:     d.Quarter4Qty1,
				Quarter4Qty2:     d.Quarter4Qty2,
				Quarter4QtyTotal: d.Quarter4QtyTotal,
			}
			rst = append(rst, item)
		}
	}

	return
}

func (s *sNurseMonthly) getIndexByYears(years float64) int {
	if years < 1 {
		return 1
	} else if years < 2 {
		return 2
	} else if years < 5 {
		return 3
	} else if years < 10 {
		return 4
	} else if years < 20 {
		return 5
	} else {
		return 6
	}
}
func (s *sNurseMonthly) addQty(d *model.QuarterTotalChange, detail []*model.QuarterTotalInfoRes, startIndex, endIndex int, qty int, curYear int) {
	if d.TransYear > curYear {
		if d.StartYear < curYear {
			detail[startIndex].Quarter1Qty1 = detail[startIndex].Quarter1Qty1 - qty
			detail[startIndex].Quarter1Qty2 = detail[startIndex].Quarter1Qty2 + qty
		} else if d.StartYear == curYear && d.StartQuarter == 1 {
			detail[endIndex].Quarter1Qty1 = detail[endIndex].Quarter1Qty1 - qty
			detail[endIndex].Quarter1Qty2 = detail[endIndex].Quarter1Qty2 + qty

			detail[startIndex].Quarter2Qty1 = detail[startIndex].Quarter2Qty1 - qty
			detail[startIndex].Quarter2Qty2 = detail[startIndex].Quarter2Qty2 + qty
		}

		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter <= 2 {
			detail[endIndex].Quarter2Qty1 = detail[endIndex].Quarter2Qty1 - qty
			detail[endIndex].Quarter2Qty2 = detail[endIndex].Quarter2Qty2 + qty

			detail[startIndex].Quarter3Qty1 = detail[startIndex].Quarter3Qty1 - qty
			detail[startIndex].Quarter3Qty2 = detail[startIndex].Quarter3Qty2 + qty
		}

		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter <= 3 {
			detail[endIndex].Quarter3Qty1 = detail[endIndex].Quarter3Qty1 - qty
			detail[endIndex].Quarter3Qty2 = detail[endIndex].Quarter3Qty2 + qty

			detail[startIndex].Quarter4Qty1 = detail[startIndex].Quarter4Qty1 - qty
			detail[startIndex].Quarter4Qty2 = detail[startIndex].Quarter4Qty2 + qty
		}
		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter == 4 {
			detail[endIndex].Quarter4Qty1 = detail[endIndex].Quarter4Qty1 - qty
			detail[endIndex].Quarter4Qty2 = detail[endIndex].Quarter4Qty2 + qty
		}
		return
	}

	if d.QuarterIndex > 1 {
		if d.StartYear < curYear {
			detail[startIndex].Quarter1Qty1 = detail[startIndex].Quarter1Qty1 - qty
			detail[startIndex].Quarter1Qty2 = detail[startIndex].Quarter1Qty2 + qty

			detail[endIndex].Quarter1Qty1 = detail[endIndex].Quarter1Qty1 - qty
			detail[endIndex].Quarter1Qty2 = detail[endIndex].Quarter1Qty2 + qty
		} else if d.StartYear == curYear && d.StartQuarter == 1 {
			detail[endIndex].Quarter1Qty1 = detail[endIndex].Quarter1Qty1 - qty
			detail[endIndex].Quarter1Qty2 = detail[endIndex].Quarter1Qty2 + qty
		}
	}
	if d.QuarterIndex == 2 {
		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter == 1 {
			detail[startIndex].Quarter2Qty1 = detail[startIndex].Quarter2Qty1 - qty
			detail[startIndex].Quarter2Qty2 = detail[startIndex].Quarter2Qty2 + qty

		}
	}

	if d.QuarterIndex > 2 {
		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter == 1 {
			detail[startIndex].Quarter2Qty1 = detail[startIndex].Quarter2Qty1 - qty
			detail[startIndex].Quarter2Qty2 = detail[startIndex].Quarter2Qty2 + qty

			detail[endIndex].Quarter2Qty1 = detail[endIndex].Quarter2Qty1 - qty
			detail[endIndex].Quarter2Qty2 = detail[endIndex].Quarter2Qty2 + qty
		} else if d.StartYear == curYear && d.StartQuarter == 2 {
			detail[endIndex].Quarter2Qty1 = detail[endIndex].Quarter2Qty1 - qty
			detail[endIndex].Quarter2Qty2 = detail[endIndex].Quarter2Qty2 + qty
		}
	}

	if d.QuarterIndex == 3 {
		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter < 3 {
			detail[startIndex].Quarter3Qty1 = detail[startIndex].Quarter3Qty1 - qty
			detail[startIndex].Quarter3Qty2 = detail[startIndex].Quarter3Qty2 + qty
		}
	}

	if d.QuarterIndex == 4 {
		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter < 3 {
			detail[startIndex].Quarter3Qty1 = detail[startIndex].Quarter3Qty1 - qty
			detail[startIndex].Quarter3Qty2 = detail[startIndex].Quarter3Qty2 + qty

			detail[endIndex].Quarter3Qty1 = detail[endIndex].Quarter3Qty1 - qty
			detail[endIndex].Quarter3Qty2 = detail[endIndex].Quarter3Qty2 + qty
		} else if d.StartYear == curYear && d.StartQuarter == 3 {
			detail[endIndex].Quarter3Qty1 = detail[endIndex].Quarter3Qty1 - qty
			detail[endIndex].Quarter3Qty2 = detail[endIndex].Quarter3Qty2 + qty
		}

		if d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter < 4 {
			detail[startIndex].Quarter4Qty1 = detail[startIndex].Quarter4Qty1 - qty
			detail[startIndex].Quarter4Qty2 = detail[startIndex].Quarter4Qty2 + qty
		}

	}
}
func (s *sNurseMonthly) addQtyQ0(d *model.QuarterTotalChange, detail []*model.QuarterTotalInfoRes, startIndex, qty int, curYear int) {

	if d.QuarterIndex > 1 && (d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter == 1) {
		detail[startIndex].Quarter1Qty1 = detail[startIndex].Quarter1Qty1 - qty
		detail[startIndex].Quarter1Qty2 = detail[startIndex].Quarter1Qty2 + qty
	}
}

func (s *sNurseMonthly) addQtyQ1(d *model.QuarterTotalChange, detail []*model.QuarterTotalInfoRes, startIndex, endIndex int, qty int, curYear int) {

	if d.QuarterIndex > 1 && (d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter == 1) {
		detail[endIndex].Quarter1Qty1 = detail[endIndex].Quarter1Qty1 - qty
		detail[endIndex].Quarter1Qty2 = detail[endIndex].Quarter1Qty2 + qty

		detail[startIndex].Quarter2Qty1 = detail[startIndex].Quarter2Qty1 - qty
		detail[startIndex].Quarter2Qty2 = detail[startIndex].Quarter2Qty2 + qty
	}
}

func (s *sNurseMonthly) addQtyQ2(d *model.QuarterTotalChange, detail []*model.QuarterTotalInfoRes, startIndex, endIndex int, qty int, curYear int) {

	if d.QuarterIndex > 2 && (d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter <= 2) {
		detail[endIndex].Quarter2Qty1 = detail[endIndex].Quarter2Qty1 - qty
		detail[endIndex].Quarter2Qty2 = detail[endIndex].Quarter2Qty2 + qty

		detail[startIndex].Quarter3Qty1 = detail[startIndex].Quarter3Qty1 - qty
		detail[startIndex].Quarter3Qty2 = detail[startIndex].Quarter3Qty2 + qty
	}

}

func (s *sNurseMonthly) addQtyQ3(d *model.QuarterTotalChange, detail []*model.QuarterTotalInfoRes, startIndex, endIndex int, qty int, curYear int) {

	if d.QuarterIndex > 3 && (d.StartYear < curYear || d.StartYear == curYear && d.StartQuarter <= 3) {

		detail[endIndex].Quarter3Qty1 = detail[endIndex].Quarter3Qty1 - qty
		detail[endIndex].Quarter3Qty2 = detail[endIndex].Quarter3Qty2 + qty

		detail[startIndex].Quarter4Qty1 = detail[startIndex].Quarter4Qty1 - qty
		detail[startIndex].Quarter4Qty2 = detail[startIndex].Quarter4Qty2 + qty
	}

}
func (s *sNurseMonthly) ListQuarterTotal(ctx context.Context, req *report.QuarterTotalSearchReq) (res *report.QuarterTotalSearchRes, err error) {
	res = new(report.QuarterTotalSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			detail    []*model.QuarterTotalInfoRes
			deptList  []*archiveEntity.Department
			titleList []*archiveEntity.Title
			eduList   []*archiveEntity.Education

			detailChange []*model.QuarterTotalChange
			ids          []string
			ids2         []string
			ids3         []string
			ids4         []string
			ids5         []string
			whereSql     string
			fieldSql     string
			groupSql     string
			sortSql      string
			fieldSql2    string

			curYear   int
			difQty    int
			lastIndex int
			curIndex  int
			tempIndex int
		)
		if req.Year == "" {
			req.Year = strconv.Itoa(time.Now().Year())
		}
		curYear, err = strconv.Atoi(req.Year)

		detail, err = s.getQuarterTotal(ctx, "护士数量配置相关数据", "全院执业护士总人数", "11",
			req.Year, "", "", "", "")
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.List = append(res.List, detail...)

		err = dao.Department.Ctx(ctx).Fields("dept_id").Where("dept_type='住院病区'").Scan(&deptList)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		for _, d := range deptList {
			ids = append(ids, strconv.FormatInt(d.DeptId, 10))
		}
		whereSql = " and dept_id in(" + strings.Join(ids, ",") + ")"
		detail, err = s.getQuarterTotal(ctx, "护士数量配置相关数据", "住院病区执业护士总人数", "12",
			req.Year, "", "", "", whereSql)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.List = append(res.List, detail...)

		ids = nil
		err = dao.Title.Ctx(ctx).Fields("title_id,title_name").Scan(&titleList)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		for _, d := range titleList {
			switch d.TitleName {
			case "N1级", "N2级":
				ids = append(ids, strconv.FormatInt(d.TitleId, 10))
			case "N3级":
				ids2 = append(ids2, strconv.FormatInt(d.TitleId, 10))
			case "N4级":
				ids3 = append(ids3, strconv.FormatInt(d.TitleId, 10))
			case "N5级":
				ids4 = append(ids4, strconv.FormatInt(d.TitleId, 10))
			case "N6级":
				ids5 = append(ids5, strconv.FormatInt(d.TitleId, 10))
			}

		}

		fieldSql = "case when title_id in(" + strings.Join(ids, ",") + ") then '护士（初级）人数' " +
			" when title_id in(" + strings.Join(ids2, ",") + ") then '护师人数' " +
			" when title_id in(" + strings.Join(ids3, ",") + ") then '主管护师人数' " +
			" when title_id in(" + strings.Join(ids4, ",") + ") then '副主任护师人数' " +
			" when title_id in(" + strings.Join(ids5, ",") + ") then '主任护师人数' else '' end as item_name," +
			"case when title_id in(" + strings.Join(ids, ",") + ") then 1 " +
			" when title_id in(" + strings.Join(ids2, ",") + ") then 2 " +
			" when title_id in(" + strings.Join(ids3, ",") + ") then 3 " +
			" when title_id in(" + strings.Join(ids4, ",") + ") then 4 " +
			" when title_id in(" + strings.Join(ids5, ",") + ") then 5 " +
			"else 0 end as item_key,"
		groupSql = "item_name"
		sortSql = "item_key"

		detail, err = s.getQuarterTotal(ctx, "人力资源结构--职称相关数据", "", "2",
			req.Year, fieldSql, groupSql, sortSql, "")
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		if len(detail) < 10 {
			lastIndex = 0
			for _, d := range detail {
				curIndex, err = strconv.Atoi(string(d.ItemKey[1]))
				if curIndex > lastIndex+1 {
					for j := lastIndex + 1; j < curIndex; j++ {
						switch j {
						case 1:
							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季初护士（初级）人数",
								ItemKey:   "211",
							})

							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季末护士（初级）人数",
								ItemKey:   "212",
							})
						case 2:
							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季初护师人数",
								ItemKey:   "221",
							})

							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季末护师人数",
								ItemKey:   "222",
							})
						case 3:
							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季初主管护师人数",
								ItemKey:   "231",
							})

							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季末主管护师人数",
								ItemKey:   "232",
							})
						case 4:
							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季初副主任护师人数",
								ItemKey:   "241",
							})

							res.List = append(res.List, &model.QuarterTotalInfoRes{
								Year:      req.Year,
								ItemClass: "人力资源结构--职称相关数据",
								ItemName:  "季末副主任护师人数",
								ItemKey:   "242",
							})
						}
					}
				}
				res.List = append(res.List, d)
				lastIndex = curIndex
			}
		} else {
			res.List = append(res.List, detail...)
		}

		fieldSql = "title_id,case"
		fieldSql2 = "case"
		err = dao.Education.Ctx(ctx).Fields("education_id,education_name,seq").Order("seq").Scan(&eduList)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		for i, d := range eduList {
			fieldSql = fieldSql + " when education_id =" + strconv.FormatInt(d.EducationId, 10) + " then '" + d.EducationName + "护士人数'"
			fieldSql2 = fieldSql2 + " when education_id =" + strconv.FormatInt(d.EducationId, 10) + " then '" + strconv.Itoa(i+1) + "'"
		}

		fieldSql = fieldSql + " else '无' end as item_name,"
		fieldSql2 = fieldSql2 + " else '' end as item_key,"
		groupSql = "item_key"
		sortSql = "item_key"

		detail, err = s.getQuarterTotal(ctx, "人力资源结构--学历相关数据", "", "3",
			req.Year, fieldSql+fieldSql2, groupSql, sortSql, "")
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		if len(detail) < len(eduList)*2 {
			lastIndex = 0
			for _, d := range detail {
				curIndex, err = strconv.Atoi(string(d.ItemKey[1]))
				if curIndex > lastIndex+1 {
					for j := lastIndex + 1; j < curIndex; j++ {
						res.List = append(res.List, &model.QuarterTotalInfoRes{
							Year:      req.Year,
							ItemClass: "人力资源结构--学历相关数据",
							ItemName:  "季初" + eduList[j-1].EducationName + "护士人数",
							ItemKey:   "3" + strconv.Itoa(j) + "1",
						})

						res.List = append(res.List, &model.QuarterTotalInfoRes{
							Year:      req.Year,
							ItemClass: "人力资源结构--学历相关数据",
							ItemName:  "季末" + eduList[j-1].EducationName + "护士人数",
							ItemKey:   "3" + strconv.Itoa(j) + "2",
						})
					}
				}
				res.List = append(res.List, d)
				lastIndex = curIndex
			}
		} else {
			res.List = append(res.List, detail...)
		}

		detail, err = s.getQuarterTotalAges(ctx, req.Year)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.List = append(res.List, detail...)

		whereSql = "d.change_type=1 and d.from_area_id<>d.to_area_id and d.trans_date>='" + req.Year + "-01-01'"
		fieldSql = "DATE_FORMAT(d.trans_date,'%Y-%m') month_str,year(d.trans_date) as trans_year,QUARTER(d.trans_date) as quarter_index," +
			"DATE_FORMAT(t_nurse.start_date,'%Y-%m') start_month,year(t_nurse.start_date) as start_year,QUARTER(t_nurse.start_date) as start_quarter," +
			"d.nurse_id,d.from_area_id,d.to_area_id,dep.dept_type,e.seq as education_seq,t.title_name," +
			"DATEDIFF('" + strconv.Itoa(curYear-1) + "-12-31',start_date)/365.0 as y0," +
			"DATEDIFF('" + req.Year + "-03-31',start_date)/365.0 as y1," +
			"DATEDIFF('" + req.Year + "-06-30',start_date)/365.0 as y2," +
			"DATEDIFF('" + req.Year + "-09-30',start_date)/365.0 as y3"
		err = dao.Nurse.Ctx(ctx).InnerJoin("t_nurse_change_log d", "d.nurse_id=t_nurse.nurse_id").
			InnerJoin("t_department dep", "dep.dept_id=t_nurse.dept_id").
			InnerJoin("t_education e", "e.education_id=t_nurse.education_id").
			InnerJoin("t_title t", "t.title_id=t_nurse.title_id").
			Fields(fieldSql).
			Where(whereSql).Scan(&detailChange)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")

		for _, d := range detailChange {
			if d.FromAreaId == 1 {
				// area_id 1-->2,Quarter1Qty1+1,  Quarter1Qty2-1
				difQty = -1
			} else {
				// area_id 2-->1,Quarter1Qty1-1,  Quarter1Qty2+1
				difQty = 1
			}
			s.addQty(d, res.List, 0, 1, difQty, curYear)
			if d.DeptType == "住院病区" {
				s.addQty(d, res.List, 2, 3, difQty, curYear)
			}
			if d.TitleName == "N1级" || d.TitleName == "N2级" {
				s.addQty(d, res.List, 4, 5, difQty, curYear)
			}

			if d.TitleName == "N3级" {
				s.addQty(d, res.List, 6, 7, difQty, curYear)
			}
			if d.TitleName == "N4级" {
				s.addQty(d, res.List, 8, 9, difQty, curYear)
			}

			if d.TitleName == "N5级" {
				s.addQty(d, res.List, 10, 11, difQty, curYear)
			}
			if d.TitleName == "N6级" {
				s.addQty(d, res.List, 12, 13, difQty, curYear)
			}

			if d.EducationSeq >= 1 && d.EducationSeq <= len(eduList) {
				s.addQty(d, res.List, 13+d.EducationSeq*2-1, 13+d.EducationSeq*2, difQty, curYear)
			}
			lastIndex = 13 + len(eduList)*2
			tempIndex = s.getIndexByYears(d.Y0)
			s.addQtyQ0(d, res.List, lastIndex+(tempIndex*2-1), difQty, curYear)

			tempIndex = s.getIndexByYears(d.Y1)
			s.addQtyQ1(d, res.List, lastIndex+(tempIndex*2-1), lastIndex+tempIndex*2, difQty, curYear)

			tempIndex = s.getIndexByYears(d.Y2)
			s.addQtyQ2(d, res.List, lastIndex+(tempIndex*2-1), lastIndex+tempIndex*2, difQty, curYear)

			tempIndex = s.getIndexByYears(d.Y3)
			s.addQtyQ3(d, res.List, lastIndex+(tempIndex*2-1), lastIndex+tempIndex*2, difQty, curYear)

		}

		detail, err = s.getQuarterLevel(ctx, "离职相关数据", "执业护士离职总人数",
			req.Year, "", "", "", "")
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.List = append(res.List, detail...)

		fieldSql = "case when title_id in(" + strings.Join(ids, ",") + ") then '护士（初级）离职人数' " +
			" when title_id in(" + strings.Join(ids2, ",") + ") then '护师离职人数' " +
			" when title_id in(" + strings.Join(ids3, ",") + ") then '主管护师离职人数' " +
			" when title_id in(" + strings.Join(ids4, ",") + ") then '副主任护师离职人数' " +
			" when title_id in(" + strings.Join(ids5, ",") + ") then '主任护师离职人数' else '' end as item_name," +
			"case when title_id in(" + strings.Join(ids, ",") + ") then 1 " +
			" when title_id in(" + strings.Join(ids2, ",") + ") then 2 " +
			" when title_id in(" + strings.Join(ids3, ",") + ") then 3 " +
			" when title_id in(" + strings.Join(ids4, ",") + ") then 4 " +
			" when title_id in(" + strings.Join(ids5, ",") + ") then 5 " +
			"else '' end as item_key,"
		groupSql = "item_name"
		sortSql = "item_key"

		detail, err = s.getQuarterLevel(ctx, "离职相关数据", "",
			req.Year, fieldSql, groupSql, sortSql, "")
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.List = append(res.List, detail...)

		fieldSql = "case when DATEDIFF(curdate(),start_date)/365.0<1 then '<1年资护士离职人数' " +
			" when DATEDIFF(curdate(),start_date)/365.0<2 then '1≤y<2年资护士离职人数' " +
			" when DATEDIFF(curdate(),start_date)/365.0<5 then '2≤y<5年资护士离职人数' " +
			" when DATEDIFF(curdate(),start_date)/365.0<10 then '5≤y<10年资护士离职人数' " +
			" when DATEDIFF(curdate(),start_date)/365.0<20 then '10≤y<20年资护士离职人数' "
		fieldSql = fieldSql + " else '20年资护士离职人数' end as item_name,"

		fieldSql = fieldSql + "case when DATEDIFF(curdate(),start_date)/365.0<1 then 0 " +
			" when DATEDIFF(curdate(),start_date)/365.0<2 then 1 " +
			" when DATEDIFF(curdate(),start_date)/365.0<5 then 2 " +
			" when DATEDIFF(curdate(),start_date)/365.0<10 then 3 " +
			" when DATEDIFF(curdate(),start_date)/365.0<20 then 4 "
		fieldSql = fieldSql + " else 5 end as item_key,"

		groupSql = "item_key"
		sortSql = "item_key"

		detail, err = s.getQuarterLevel(ctx, "离职相关数据", "",
			req.Year, fieldSql, groupSql, sortSql, "")
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.List = append(res.List, detail...)

		/*
			curItemClass = "护士数量配置相关数据"
			res.AppendItem(req.Year, curItemClass, "季初全院执业护士总人数", "")
			res.AppendItem(req.Year, curItemClass, "季末全院执业护士总人数", "")
			res.AppendItem(req.Year, curItemClass, "季初住院病区执业护士总人数", "")
			res.AppendItem(req.Year, curItemClass, "季末住院病区执业护士总人数", "")

			curItemClass = "人力资源结构--职称相关数据"
			res.AppendItem(req.Year, curItemClass, "季初护士（初级）人数", "")
			res.AppendItem(req.Year, curItemClass, "季末护士（初级）人数", "")
			res.AppendItem(req.Year, curItemClass, "季初护师人数", "")
			res.AppendItem(req.Year, curItemClass, "季末护师人数", "")
			res.AppendItem(req.Year, curItemClass, "季初主管护师人数", "")
			res.AppendItem(req.Year, curItemClass, "季末主管护师人数", "")

			res.AppendItem(req.Year, curItemClass, "季初副主任护师人数", "")
			res.AppendItem(req.Year, curItemClass, "季末副主任护师人数", "")
			res.AppendItem(req.Year, curItemClass, "季初主任护师人数", "")
			res.AppendItem(req.Year, curItemClass, "季末主任护师人数", "")

			curItemClass = "人力资源结构--学历相关数据"
			res.AppendItem(req.Year, curItemClass, "季初中专护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末中专护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初大专护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末大专护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初本科护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末本科护士人数", "")

			res.AppendItem(req.Year, curItemClass, "季初硕士护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末硕士护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初博士护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末博士护士人数", "")

			curItemClass = "人力资源结构--工作年限相关数据"
			res.AppendItem(req.Year, curItemClass, "季初<1年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末<1年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初1≤y<2年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末1≤y<2年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初2≤y<5年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末2≤y<5年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初5≤y<10年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末5≤y<10年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季初10≤y<20年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末10≤y<20年资护士人数", "")

			res.AppendItem(req.Year, curItemClass, "季初≥20年资护士人数", "")
			res.AppendItem(req.Year, curItemClass, "季末≥20年资护士人数", "")

			curItemClass = "离职相关数据"
			res.AppendItem(req.Year, curItemClass, "执业护士离职总人数", "")
			res.AppendItem(req.Year, curItemClass, "护士（初级）离职人数", "")
			res.AppendItem(req.Year, curItemClass, "护师离职人数", "")
			res.AppendItem(req.Year, curItemClass, "主管护师离职人数", "")
			res.AppendItem(req.Year, curItemClass, "副主任护师离职人数", "")
			res.AppendItem(req.Year, curItemClass, "主任护师离职人数", "")
			res.AppendItem(req.Year, curItemClass, "中专护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "大专护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "本科护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "硕士护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "博士护士离职人数", "")

			res.AppendItem(req.Year, curItemClass, "<1年资护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "1≤y<2年资护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "2≤y<5年资护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "5≤y<10年资护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "10≤y<20年资护士离职人数", "")
			res.AppendItem(req.Year, curItemClass, "≥20年资护士离职人数", "")
		*/
	})
	return
}

func (s *sNurseMonthly) GetHomeData(ctx context.Context, req *report.HomeSearchReq) (res *report.HomeSearchRes, err error) {
	res = new(report.HomeSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			detail        []*model.MonthlyTotalInfoRes
			areaList      []*commonEntity.SysDictData
			total         int
			monthTotalReq report.MonthlyTotalSearchReq
			monthTotalRes *report.MonthlyTotalSearchRes
			eduDetail     []report.HomeSearchEduGroupRes
			items         []int
			leaveItems    []int
		)

		err = commonDao.SysDictData.Ctx(ctx).Fields(commonDao.SysDictData.Columns().DictLabel, commonDao.SysDictData.Columns().DictValue).
			Where(commonDao.SysDictData.Columns().DictType+" = ?", "area_id").Order(commonDao.SysDictData.Columns().DictValue).Scan(&areaList)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")

		err = dao.Nurse.Ctx(ctx).Fields("area_id,count(1) as add_qty").Where("status=1").
			Group("area_id").Order("area_id").Scan(&detail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		total = 0
		for i, d := range detail {
			for _, area := range areaList {
				if strconv.Itoa(d.AreaId) == area.DictValue {
					detail[i].MonthStr = area.DictLabel
				}
			}
			total = total + d.AddQty
		}

		res.TotalData = append(res.TotalData, &report.HomeSearchTotalRes{
			ItemName:   "总人数",
			ItemValue:  strconv.Itoa(total),
			ItemValue2: "100.00",
		})

		for _, d := range detail {
			item := &report.HomeSearchTotalRes{
				ItemName:  d.MonthStr,
				ItemValue: strconv.Itoa(d.AddQty),
			}
			if total != 0 {
				item.ItemValue2 = fmt.Sprintf("%.2f", float64(d.AddQty)/float64(total)*100.0)
			}

			res.TotalData = append(res.TotalData, item)

		}
		monthTotalReq = report.MonthlyTotalSearchReq{}
		if int(time.Now().Month()) == 1 {
			monthTotalReq.Year = strconv.Itoa(time.Now().Year() - 1)
		} else {
			monthTotalReq.Year = strconv.Itoa(time.Now().Year())
		}

		monthTotalRes, err = s.ListTotal(ctx, &monthTotalReq)

		for _, area := range areaList {
			res.AreaData.AreaList = append(res.AreaData.AreaList, area.DictLabel)
			items = []int{}
			leaveItems = []int{}
			for _, d := range monthTotalRes.List {
				if area.DictValue == strconv.Itoa(d.AreaId) {
					items = append(items, d.EndQty)
					leaveItems = append(leaveItems, d.LeaveQty)
				}
			}
			res.AreaData.Data = append(res.AreaData.Data, items)
			res.AreaData.LeaveData = append(res.AreaData.LeaveData, leaveItems)
		}
		err = dao.Nurse.Ctx(ctx).LeftJoin("t_education e ", "e.education_id=t_nurse.education_id").
			Fields("e.education_name as itemName,count(1) as itemValue").Where("t_nurse.status=1").
			Group("e.education_name").Order("e.seq").Scan(&eduDetail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.EduData = append(res.EduData, eduDetail...)

		err = dao.Nurse.Ctx(ctx).LeftJoin("t_title e ", "e.title_id=t_nurse.title_id").
			Fields("e.title_name as itemName,count(1) as itemValue").Where("t_nurse.status=1").
			Group("e.title_name").Order("e.seq").Scan(&eduDetail)
		liberr.ErrIsNilEx(ctx, err, "获取数据错误")
		res.TitleData = append(res.TitleData, eduDetail...)
	})
	return
}
