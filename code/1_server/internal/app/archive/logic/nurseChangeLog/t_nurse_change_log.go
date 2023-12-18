// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/logic/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tiger1103/gfast/v3/api/v1/archive"
	"github.com/tiger1103/gfast/v3/internal/app/archive/dao"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model"
	"github.com/tiger1103/gfast/v3/internal/app/archive/model/do"
	"github.com/tiger1103/gfast/v3/internal/app/archive/service"
	"github.com/tiger1103/gfast/v3/internal/app/system/consts"
	systemService "github.com/tiger1103/gfast/v3/internal/app/system/service"
	"github.com/tiger1103/gfast/v3/library/liberr"
)

func init() {
	service.RegisterNurseChangeLog(New())
}

func New() *sNurseChangeLog {
	return &sNurseChangeLog{}
}

type sNurseChangeLog struct{}

func (s *sNurseChangeLog) List(ctx context.Context, req *archive.NurseChangeLogSearchReq) (listRes *archive.NurseChangeLogSearchRes, err error) {
	listRes = new(archive.NurseChangeLogSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NurseChangeLog.Ctx(ctx).WithAll()
		if req.NurseId != "" {
			m = m.Where(dao.NurseChangeLog.Columns().NurseId+" = ?", gconv.Int64(req.NurseId))
		}
		if req.ChangeType != "" {
			m = m.Where(dao.NurseChangeLog.Columns().ChangeType+" = ?", gconv.Int(req.ChangeType))
		}
		if req.FromAreaId != "" {
			m = m.Where(dao.NurseChangeLog.Columns().FromAreaId+" = ?", gconv.Int(req.FromAreaId))
		}
		if req.FromDeptId != "" {
			m = m.Where(dao.NurseChangeLog.Columns().FromDeptId+" = ?", gconv.Int64(req.FromDeptId))
		}
		if req.ToAreaId != "" {
			m = m.Where(dao.NurseChangeLog.Columns().ToAreaId+" = ?", gconv.Int(req.ToAreaId))
		}
		if req.ToDeptId != "" {
			m = m.Where(dao.NurseChangeLog.Columns().ToDeptId+" = ?", gconv.Int64(req.ToDeptId))
		}
		if req.TransDate != "" {
			m = m.Where(dao.NurseChangeLog.Columns().TransDate+" = ?", gconv.Time(req.TransDate))
		}
		if req.Status != "" {
			m = m.Where(dao.NurseChangeLog.Columns().Status+" = ?", gconv.Int(req.Status))
		}
		if len(req.DateRange) != 0 {
			m = m.Where(dao.NurseChangeLog.Columns().CreatedAt+" >=? AND "+dao.NurseChangeLog.Columns().CreatedAt+" <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.CreatedBy != "" {
			m = m.Where(dao.NurseChangeLog.Columns().CreatedBy+" = ?", req.CreatedBy)
		}
		if req.DeletedBy != "" {
			m = m.Where(dao.NurseChangeLog.Columns().DeletedBy+" = ?", req.DeletedBy)
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
		order := "id asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.NurseChangeLogInfoRes
		err = m.Fields(archive.NurseChangeLogSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.NurseChangeLogListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.NurseChangeLogListRes{
				Id:             v.Id,
				NurseId:        v.NurseId,
				LinkedNurseId:  v.LinkedNurseId,
				ChangeType:     v.ChangeType,
				FromEntityName: v.FromEntityName,
				ToEntityName:   v.ToEntityName,
				TransDate:      v.TransDate,
				CreatedAt:      v.CreatedAt,
				CreatedBy:      v.CreatedBy,
			}
		}
	})
	return
}

func (s *sNurseChangeLog) GetById(ctx context.Context, id int64) (res *model.NurseChangeLogInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.NurseChangeLog.Ctx(ctx).WithAll().Where(dao.NurseChangeLog.Columns().Id, id).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sNurseChangeLog) Add(ctx context.Context, req *archive.NurseChangeLogAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.NurseChangeLog.Ctx(ctx).Insert(do.NurseChangeLog{
			NurseId:    req.NurseId,
			ChangeType: req.ChangeType,
			FromAreaId: req.FromAreaId,
			FromDeptId: req.FromDeptId,
			ToAreaId:   req.ToAreaId,
			ToDeptId:   req.ToDeptId,
			TransDate:  req.TransDate,
			Status:     req.Status,
			DeletedBy:  req.DeletedBy,
			CreatedBy:  systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sNurseChangeLog) Edit(ctx context.Context, req *archive.NurseChangeLogEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.NurseChangeLog.Ctx(ctx).WherePri(req.Id).Update(do.NurseChangeLog{
			NurseId:    req.NurseId,
			ChangeType: req.ChangeType,
			FromAreaId: req.FromAreaId,
			FromDeptId: req.FromDeptId,
			ToAreaId:   req.ToAreaId,
			ToDeptId:   req.ToDeptId,
			TransDate:  req.TransDate,
			Status:     req.Status,
			DeletedBy:  req.DeletedBy,
			UpdatedBy:  systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sNurseChangeLog) Delete(ctx context.Context, ids []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var ids []int64
		ids = append(ids, ids...)
		_, err = dao.NurseChangeLog.Ctx(ctx).Delete(dao.NurseChangeLog.Columns().Id+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
