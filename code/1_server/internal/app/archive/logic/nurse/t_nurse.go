// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-12-01 20:01:44
// 生成路径: internal/app/archive/logic/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/dao"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model/entity"
	"github.com/tiger1103/gfast/v3/internal/app/archive/service"
	commonDao "github.com/tiger1103/gfast/v3/internal/app/common/dao"
	commonModel "github.com/tiger1103/gfast/v3/internal/app/common/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
	"strings"
)

func init() {
	service.RegisterNurse(New())
}

func New() *sNurse {
	return &sNurse{}
}

type sNurse struct{}

func (s *sNurse) List(ctx context.Context, req *archive.NurseSearchReq) (listRes *archive.NurseSearchRes, err error) {
	listRes = new(archive.NurseSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Nurse.Ctx(ctx).WithAll()
		if req.AreaId != "" {
			m = m.Where(dao.Nurse.Columns().AreaId+" = ?", gconv.Int(req.AreaId))
		}
		if req.NurseName != "" {
			if strings.Index(req.NurseName, ",") > 0 {
				names := strings.Split(req.NurseName, ",")
				m = m.Where(dao.Nurse.Columns().NurseName+" in(?)", names)
			} else if strings.Index(req.NurseName, "、") > 0 {
				names := strings.Split(req.NurseName, "、")
				m = m.Where(dao.Nurse.Columns().NurseName+" in(?)", names)
			} else {
				m = m.Where(dao.Nurse.Columns().NurseName+" like ?", "%"+req.NurseName+"%")
			}
		}
		if req.DeptId != "" {
			m = m.Where(dao.Nurse.Columns().DeptId+" = ?", gconv.Int64(req.DeptId))
		}
		if req.Status != "" {
			m = m.Where(dao.Nurse.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if req.StartDate != "" {
			m = m.Where(dao.Nurse.Columns().StartDate+" = ?", gconv.Time(req.StartDate))
		}
		if req.EducationId != "" {
			m = m.Where(dao.Nurse.Columns().EducationId+" = ?", gconv.Int64(req.EducationId))
		}
		if req.TitleId != "" {
			m = m.Where(dao.Nurse.Columns().TitleId+" = ?", gconv.Int64(req.TitleId))
		}
		if req.PositionId != "" {
			m = m.Where(dao.Nurse.Columns().PositionId+" = ?", gconv.Int64(req.PositionId))
		}
		if req.StaffType != "" {
			m = m.Where(dao.Nurse.Columns().StaffType+" = ?", gconv.Int(req.StaffType))
		}
		if req.NurseCode != "" {
			m = m.Where(dao.Nurse.Columns().NurseCode+" = ?", req.NurseCode)
		}
		if req.RegStatus == "1" {
			m = m.Where(dao.Nurse.Columns().RegDate + " is not null ")
		} else if req.RegStatus == "0" {
			m = m.Where(dao.Nurse.Columns().RegDate + " is null ")
		}

		listRes.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "nurse_code asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.NurseInfoRes
		err = m.Fields(archive.NurseSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.NurseListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.NurseListRes{
				NurseId:           v.NurseId,
				Status:            v.Status,
				NurseCode:         v.NurseCode,
				NurseName:         v.NurseName,
				Sex:               v.Sex,
				Birthday:          v.Birthday,
				DeptId:            v.DeptId,
				LinkedDeptId:      v.LinkedDeptId,
				LinkedPositionId:  v.LinkedPositionId,
				StartDate:         v.StartDate,
				EducationId:       v.EducationId,
				LinkedEducationId: v.LinkedEducationId,
				TitleId:           v.TitleId,
				LinkedTitleId:     v.LinkedTitleId,
				StaffType:         v.StaffType,
				EndDate:           v.EndDate,
				Note:              v.Note,
				AreaId:            v.AreaId,
				CertNo:            v.CertNo,
				CertEnddate:       v.CertEnddate,
				RegDate:           v.RegDate,
			}
		}
	})
	return
}

func (s *sNurse) GetByNurseId(ctx context.Context, nurseId int64) (res *model.NurseInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Nurse.Ctx(ctx).WithAll().Where(dao.Nurse.Columns().NurseId, nurseId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sNurse) Add(ctx context.Context, req *archive.NurseAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Nurse.Ctx(ctx).Insert(do.Nurse{
			NurseCode:   req.NurseCode,
			NurseName:   req.NurseName,
			Sex:         req.Sex,
			Birthday:    req.Birthday,
			AreaId:      req.AreaId,
			DeptId:      req.DeptId,
			StartDate:   req.StartDate,
			EducationId: req.EducationId,
			TitleId:     req.TitleId,
			StaffType:   req.StaffType,
			Note:        req.Note,
			EndDate:     req.EndDate,
			CertNo:      req.CertNo,
			CertEnddate: req.CertEnddate,
			RegDate:     req.RegDate,
			Status:      req.Status,
			CreatedBy:   systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sNurse) Edit(ctx context.Context, req *archive.NurseEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			nurseInfo model.NurseInfoRes
			nurseLogs []*entity.NurseChangeLog
			//dictInfo  []*commonEntity.SysDictData
			dictRes   commonModel.DictDataRes
			deptInfo  model.DepartmentInfoRes
			eduInfo   model.EducationInfoRes
			titleInfo model.TitleInfoRes
			userName  string
		)
		user := systemService.Context().GetLoginUser(ctx)
		if user != nil {
			userName = user.UserName
		}

		err = dao.Nurse.Ctx(ctx).WithAll().Where(dao.Nurse.Columns().NurseId, req.NurseId).Scan(&nurseInfo)

		//err = dao.Nurse.Ctx(ctx).Fields(model.NurseInfoRes{}).WherePri(req.NurseId).Scan(&nurseInfo)
		liberr.ErrIsNil(ctx, err, "修改失败")

		if nurseInfo.AreaId != req.AreaId || nurseInfo.DeptId != req.DeptId {
			nurseLog := &entity.NurseChangeLog{
				NurseId:    req.NurseId,
				ChangeType: 1,
				FromAreaId: nurseInfo.AreaId,
				FromDeptId: nurseInfo.DeptId,
				ToAreaId:   req.AreaId,
				ToDeptId:   req.DeptId,
				TransDate:  gtime.Now(),
				CreatedBy:  userName,
				UpdatedBy:  userName,
			}
			err = commonDao.SysDictData.Ctx(ctx).Fields(commonModel.DictDataRes{}).
				Where("dict_type='area_id' and dict_value=?", nurseInfo.AreaId).Scan(&dictRes)
			liberr.ErrIsNil(ctx, err, "获取信息失败")
			nurseLog.FromEntityName = dictRes.DictLabel + ":" + nurseInfo.LinkedDeptId.DeptName

			if nurseInfo.AreaId != req.AreaId {
				err = commonDao.SysDictData.Ctx(ctx).Fields(commonModel.DictDataRes{}).
					Where("dict_type='area_id' and dict_value=?", req.AreaId).Scan(&dictRes)
				liberr.ErrIsNil(ctx, err, "获取信息失败")
			}

			err = dao.Department.Ctx(ctx).Fields(dao.Department.Columns().DeptName).WherePri(req.DeptId).Scan(&deptInfo)
			liberr.ErrIsNil(ctx, err, "获取信息失败")
			nurseLog.ToEntityName = dictRes.DictLabel + ":" + deptInfo.DeptName

			nurseLogs = append(nurseLogs, nurseLog)
		}
		if nurseInfo.TitleId != req.TitleId {
			nurseLog := &entity.NurseChangeLog{
				NurseId:        req.NurseId,
				ChangeType:     2,
				FromEntityId:   nurseInfo.TitleId,
				ToEntityId:     req.TitleId,
				TransDate:      gtime.Now(),
				CreatedBy:      userName,
				UpdatedBy:      userName,
				FromEntityName: nurseInfo.LinkedTitleId.TitleName,
			}

			err = dao.Title.Ctx(ctx).Fields(dao.Title.Columns().TitleName).WherePri(req.TitleId).Scan(&titleInfo)
			liberr.ErrIsNil(ctx, err, "获取信息失败")
			nurseLog.ToEntityName = titleInfo.TitleName

			nurseLogs = append(nurseLogs, nurseLog)
		}

		if nurseInfo.EducationId != req.EducationId {
			nurseLog := &entity.NurseChangeLog{
				NurseId:        req.NurseId,
				ChangeType:     3,
				FromEntityId:   nurseInfo.EducationId,
				FromEntityName: nurseInfo.LinkedEducationId.EducationName,
				ToEntityId:     req.EducationId,
				TransDate:      gtime.Now(),
				CreatedBy:      userName,
				UpdatedBy:      userName,
			}

			err = dao.Education.Ctx(ctx).Fields(dao.Education.Columns().EducationName).WherePri(req.EducationId).Scan(&eduInfo)
			liberr.ErrIsNil(ctx, err, "获取信息失败")
			nurseLog.ToEntityName = eduInfo.EducationName

			nurseLogs = append(nurseLogs, nurseLog)
		}
		if req.EndDate == nil {
			req.Status = 1
		} else {
			req.Status = 2
		}

		tx, err1 := g.DB().Begin(ctx)
		liberr.ErrIsNil(ctx, err1, "修改失败")

		defer func() {
			if err != nil {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()

		_, err = dao.Nurse.Ctx(ctx).TX(tx).Data(g.Map{
			dao.Nurse.Columns().NurseCode:   req.NurseCode,
			dao.Nurse.Columns().NurseName:   req.NurseName,
			dao.Nurse.Columns().Sex:         req.Sex,
			dao.Nurse.Columns().Birthday:    req.Birthday,
			dao.Nurse.Columns().AreaId:      req.AreaId,
			dao.Nurse.Columns().DeptId:      req.DeptId,
			dao.Nurse.Columns().StartDate:   req.StartDate,
			dao.Nurse.Columns().EducationId: req.EducationId,
			dao.Nurse.Columns().PositionId:  req.PositionId,
			dao.Nurse.Columns().TitleId:     req.TitleId,
			dao.Nurse.Columns().StaffType:   req.StaffType,
			dao.Nurse.Columns().Note:        req.Note,
			dao.Nurse.Columns().EndDate:     req.EndDate,
			dao.Nurse.Columns().IdNo:        req.IdNo,
			dao.Nurse.Columns().CertNo:      req.CertNo,
			dao.Nurse.Columns().CertEnddate: req.CertEndDate,
			dao.Nurse.Columns().RegDate:     req.RegDate,
			dao.Nurse.Columns().Status:      req.Status,
		}).WherePri(req.NurseId).Update()
		liberr.ErrIsNil(ctx, err, "保存失败")
		if len(nurseLogs) > 0 {
			_, err = dao.NurseChangeLog.Ctx(ctx).Save(&nurseLogs)
			liberr.ErrIsNil(ctx, err, "保存失败")
		}

		//_, err = dao.Nurse.Ctx(ctx).WherePri(req.NurseId).Update(do.Nurse{
		//	NurseCode:   req.NurseCode,
		//	NurseName:   req.NurseName,
		//	Sex:         req.Sex,
		//	Birthday:    req.Birthday,
		//	AreaId:      req.AreaId,
		//	DeptId:      req.DeptId,
		//	StartDate:   req.StartDate,
		//	EducationId: req.EducationId,
		//	TitleId:     req.TitleId,
		//	StaffType:   req.StaffType,
		//	Note:        req.Note,
		//	EndDate:     req.EndDate,
		//	CertNo:      req.CertNo,
		//	CertEndDate: req.CertEndDate,
		//	RegDate:     req.RegDate,
		//	Status:      req.Status,
		//	UpdatedBy:   systemService.Context().GetUserId(ctx),
		//})
		//

	})
	return
}

func (s *sNurse) Delete(ctx context.Context, nurseIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var ids []int64
		ids = append(ids, nurseIds...)
		_, err = dao.Nurse.Ctx(ctx).Delete(dao.Nurse.Columns().NurseId+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

func (s *sNurse) Import(ctx context.Context, req *archive.NurseImportReq) (res *archive.NurseImportRes, err error) {
	res = &archive.NurseImportRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			sqlStr string
			//sessionId     string
			importAdapter model.ImportNurse
			insertData    interface{}
			checkData     []entity.NurseImport
			existsCount   int
		)
		if len(req.Data) < 2 {
			panic("没有导入数据")
		}
		if req.AreaId < 1 {
			panic("没有选择院区")
		}
		importAdapter = model.ImportNurse{}
		importAdapter.Init(true)
		//importAdapter.SetColumn(req.Data[0])
		rst := importAdapter.LoadData(req.Data)
		if rst.ErrorCount > 0 {
			res.ErrorCount = rst.ErrorCount
			res.ErrorData = rst.ErrorData
			//res.ErrorData = append(res.ErrorData, rst.ErrorData...)
			return
		}

		for _, detail := range importAdapter.Data {
			insertData = detail
			_, err1 := g.DB().Insert(ctx, "t_nurse_import", insertData)
			liberr.ErrIsNilEx(ctx, err1, "数据处理失败")
		}
		sqlStr = "update t_nurse_import d inner join t_department a on d.dept_code=a.dept_name " +
			" set d.dept_id=a.dept_id where d.session_id=? "
		_, err1 := g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_department a on d.dept_code=a.dept_code " +
			" set d.dept_id=a.dept_id where d.session_id=? and d.dept_id is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_title a on d.title_code=a.title_name " +
			" set d.title_id=a.title_id  where d.session_id=? "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_education a on d.education_code=a.education_name " +
			" set d.education_id=a.education_id  where d.session_id=? "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d  set d.sex=1  where d.session_id=? and d.sex='男'"
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d  set d.sex=2  where d.session_id=? and d.sex='女'"
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		err1 = dao.NurseImport.Ctx(ctx).Fields("nurse_name,dept_id,title_id,education_id,seq").
			Where("session_id=? and (dept_id is null or title_id is null or education_id is null)", importAdapter.SessionId).Scan(&checkData)
		liberr.ErrIsNilEx(ctx, err1, "数据校验失败")

		if len(checkData) > 0 {
			for _, detail := range checkData {
				if detail.TitleId == 0 {
					res.ErrorData = append(res.ErrorData, &model.ImportErrorItem{
						Index:    detail.Seq,
						ErrorMsg: "职称错误",
					})
					res.ErrorCount = res.ErrorCount + 1
				} else if detail.EducationId == 0 {
					res.ErrorData = append(res.ErrorData, &model.ImportErrorItem{
						Index:    detail.Seq,
						ErrorMsg: "学历错误",
					})
					res.ErrorCount = res.ErrorCount + 1
				} else if detail.DeptId == 0 {
					res.ErrorData = append(res.ErrorData, &model.ImportErrorItem{
						Index:    detail.Seq,
						ErrorMsg: "科室错误",
					})
					res.ErrorCount = res.ErrorCount + 1
				}
			}
			return
		}

		existsCount, err1 = dao.NurseImport.Ctx(ctx).Count("session_id=? and exists(select 1 from t_nurse t "+
			"where t_nurse_import.nurse_name=t.nurse_name and t.dept_id=t_nurse_import.dept_id and t.deleted_at is null )", importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据校验失败")

		sqlStr = "insert into t_nurse(area_id,nurse_code,nurse_name,sex,birthday,dept_id,education_id,title_id," +
			"staff_type,start_date,end_date,id_no,reg_date,cert_no,cert_enddate,status,session_id,created_at,created_by,updated_at,updated_by)" +
			"select ?,nurse_code,nurse_name,sex,birthday,dept_id,education_id,title_id,staff_type,start_date,end_date,id_no,reg_date,cert_no,cert_enddate," +
			"status,session_id,created_at,created_by,created_at,created_by from t_nurse_import d where session_id=? " +
			" and not exists(select 1 from t_nurse t where d.nurse_name=t.nurse_name and t.dept_id=d.dept_id and t.deleted_at is null )"
		_, err1 = g.DB().Exec(ctx, sqlStr, req.AreaId, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "插入失败")

		sqlStr = "update t_nurse set nurse_code=800000+nurse_id  where session_id=? and nurse_code is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "插入失败")

		sqlStr = "delete from t_nurse_import where session_id=? "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "插入失败")

		res.ExistsCount = existsCount
		res.SuccessCount = len(req.Data) - 1 - existsCount
	})
	return
}

func (s *sNurse) UpdateByImport(ctx context.Context, req *archive.NurseUpdateByImportReq) (res *archive.NurseImportRes, err error) {
	res = &archive.NurseImportRes{}
	err = g.Try(ctx, func(ctx context.Context) {
		var (
			sqlStr string
			//sessionId     string
			importAdapter model.ImportNurse
			insertData    interface{}
			checkData     []entity.NurseImport
		)
		if len(req.Data) < 2 {
			panic("没有导入数据")
		}
		if req.AreaId < 1 {
			panic("没有选择院区")
		}
		importAdapter = model.ImportNurse{}
		importAdapter.Init(false)
		//importAdapter.SetColumn(req.Data[0])
		rst := importAdapter.LoadData(req.Data)
		if rst.ErrorCount > 0 {
			res.ErrorCount = rst.ErrorCount
			res.ErrorData = rst.ErrorData
			//res.ErrorData = append(res.ErrorData, rst.ErrorData...)
			return
		}

		for _, detail := range importAdapter.Data {
			insertData = detail
			_, err1 := g.DB().Insert(ctx, "t_nurse_import", insertData)
			liberr.ErrIsNilEx(ctx, err1, "数据处理失败")
		}
		sqlStr = "update t_nurse_import d inner join t_department a on d.dept_code=a.dept_name " +
			" set d.dept_id=a.dept_id where d.session_id=? "
		_, err1 := g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_department a on d.dept_code=a.dept_code " +
			" set d.dept_id=a.dept_id where d.session_id=? and d.dept_id is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_title a on d.title_code=a.title_name " +
			" set d.title_id=a.title_id  where d.session_id=? "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_education a on d.education_code=a.education_name " +
			" set d.education_id=a.education_id  where d.session_id=? "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join (select dict_label,dict_value from sys_dict_data where dict_type='staff_type' ) a" +
			" on d.staff_type_code=a.dict_label " +
			" set d.staff_type=a.dict_value  where d.session_id=? "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d  set d.sex=1  where d.session_id=? and d.sex='男'"
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d  set d.sex=2  where d.session_id=? and d.sex='女'"
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.dept_id=a.dept_id and d.nurse_name =a.nurse_name and a.area_id=?" +
			" set d.nurse_id=a.nurse_id  where d.session_id=? and a.deleted_at is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, req.AreaId, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_name =a.nurse_name and a.area_id=?" +
			" set d.nurse_id=a.nurse_id  where d.session_id=? and d.dept_id is null and d.nurse_id is null and a.deleted_at is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, req.AreaId, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_name =a.nurse_name and d.nurse_id<>a.nurse_id and a.area_id=? " +
			" set d.nurse_id=null  where d.session_id=? and d.dept_id is null and a.deleted_at is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, req.AreaId, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_name =a.nurse_name and d.birthday=a.birthday and a.area_id=?" +
			" set d.nurse_id=a.nurse_id  where d.session_id=? and d.nurse_id is null and a.deleted_at is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, req.AreaId, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_name =a.nurse_name and d.birthday=a.birthday and d.nurse_id<>a.nurse_id and a.area_id=? " +
			" set d.nurse_id=null  where d.session_id=? and d.dept_id is null and a.deleted_at is null "
		_, err1 = g.DB().Exec(ctx, sqlStr, req.AreaId, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		err1 = dao.NurseImport.Ctx(ctx).Fields("nurse_name,seq").
			Where("session_id=? and  nurse_id is null ", importAdapter.SessionId).Scan(&checkData)
		liberr.ErrIsNilEx(ctx, err1, "数据校验失败")

		if len(checkData) > 0 {
			for _, detail := range checkData {

				res.ErrorData = append(res.ErrorData, &model.ImportErrorItem{
					Index:    detail.Seq,
					ErrorMsg: "护士错误:" + detail.NurseName,
				})
				res.ErrorCount = res.ErrorCount + 1

			}
			return
		}

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_id =a.nurse_id" +
			" set a.end_date=d.end_date,a.status=d.status  where d.session_id=? and d.end_date is not null and d.status=2 "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_id =a.nurse_id" +
			" set a.id_no=d.id_no  where d.session_id=? and d.id_no<>'' "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_id =a.nurse_id" +
			" set a.cert_no=d.cert_no  where d.session_id=? and d.cert_no<>'' "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_id =a.nurse_id" +
			" set a.cert_end_date=d.cert_end_date  where d.session_id=? and d.cert_end_date is not null "
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_id =a.nurse_id" +
			" set a.reg_date=d.reg_date  where d.session_id=? and d.reg_date is not null"
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		sqlStr = "update t_nurse_import d inner join t_nurse a on d.nurse_id =a.nurse_id" +
			" set a.staff_type=d.staff_type  where d.session_id=? and d.staff_type is not null"
		_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		liberr.ErrIsNilEx(ctx, err1, "数据处理失败")

		//sqlStr = "delete from t_nurse_import where session_id=? "
		//_, err1 = g.DB().Exec(ctx, sqlStr, importAdapter.SessionId)
		//liberr.ErrIsNilEx(ctx, err1, "插入失败")

		res.SuccessCount = len(req.Data) - 1 - res.ErrorCount
	})
	return
}
