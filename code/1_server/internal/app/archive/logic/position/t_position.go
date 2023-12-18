// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/logic/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================

package logic

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"

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
	service.RegisterPosition(New())
}

func New() *sPosition {
	return &sPosition{}
}

type sPosition struct{}

func (s *sPosition) List(ctx context.Context, req *archive.PositionSearchReq) (listRes *archive.PositionSearchRes, err error) {
	listRes = new(archive.PositionSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Position.Ctx(ctx).WithAll()
		if req.PositionName != "" {
			m = m.Where(dao.Position.Columns().PositionName+" like ?", "%"+req.PositionName+"%")
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
		order := "seq asc"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.PositionInfoRes
		err = m.Fields(archive.PositionSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.PositionListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.PositionListRes{
				PositionId:   v.PositionId,
				PositionName: v.PositionName,
				Seq:          v.Seq,
				Status:       v.Status,
			}
		}
	})
	return
}

func (s *sPosition) GetByPositionId(ctx context.Context, positionId int64) (res *model.PositionInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Position.Ctx(ctx).WithAll().Where(dao.Position.Columns().PositionId, positionId).Scan(&res)
		liberr.ErrIsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sPosition) Add(ctx context.Context, req *archive.PositionAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Position.Ctx(ctx).Insert(do.Position{
			PositionName: req.PositionName,
			Seq:          req.Seq,
			Status:       req.Status,
			CreatedBy:    systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加失败")
	})
	return
}

func (s *sPosition) Edit(ctx context.Context, req *archive.PositionEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Position.Ctx(ctx).WherePri(req.PositionId).Update(do.Position{
			PositionName: req.PositionName,
			Seq:          req.Seq,
			Status:       req.Status,
			UpdatedBy:    systemService.Context().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改失败")
	})
	return
}

func (s *sPosition) Delete(ctx context.Context, positionIds []int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var ids []int64
		ids = append(ids, positionIds...)
		_, err = dao.Position.Ctx(ctx).Delete(dao.Position.Columns().PositionId+" in (?)", ids)
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}
