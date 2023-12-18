// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-11-30 18:32:51
// 生成路径: internal/app/archive/logic/t_title.go
// 生成人：xiao
// desc:职称
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
	service.RegisterTitle(New())
}


func New() *sTitle {
	return &sTitle{}
}


type sTitle struct{}


func (s *sTitle)List(ctx context.Context, req *archive.TitleSearchReq) (listRes *archive.TitleSearchRes, err error){
    listRes = new(archive.TitleSearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.Title.Ctx(ctx).WithAll()        
        if req.TitleName != "" {
            m = m.Where(dao.Title.Columns().TitleName+" like ?", "%"+req.TitleName+"%")
        }          
        if req.Status != "" {
            m = m.Where(dao.Title.Columns().Status+" = ?", gconv.Int(req.Status))
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
        order:= "title_id asc"
        if req.OrderBy!=""{
            order = req.OrderBy
        }
        var res []*model.TitleInfoRes
        err = m.Fields(archive.TitleSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)        
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.TitleListRes,len(res))
        for k,v:=range res{            
            listRes.List[k] = &model.TitleListRes{                  
                TitleId : v.TitleId,                    
                Seq : v.Seq,                    
                TitleName : v.TitleName,                    
                Status : v.Status,                    
                CreatedAt : v.CreatedAt,                    
                CreatedBy : v.CreatedBy,                
            }
        }
    })
    return
}


func (s *sTitle)GetByTitleId(ctx context.Context, titleId int64) (res *model.TitleInfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.Title.Ctx(ctx).WithAll().Where(dao.Title.Columns().TitleId,  titleId).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sTitle)Add(ctx context.Context, req *archive.TitleAddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {        
        _, err = dao.Title.Ctx(ctx).Insert(do.Title{            
            TitleName:req.TitleName,            
            Seq:req.Seq,            
            Status:req.Status,            
            CreatedBy:systemService.Context().GetUserId(ctx),            
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sTitle)Edit(ctx context.Context, req *archive.TitleEditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {        
        _, err = dao.Title.Ctx(ctx).WherePri(req.TitleId).Update(do.Title{            
            TitleName:req.TitleName,            
            Seq:req.Seq,            
            Status:req.Status,            
            UpdatedBy:systemService.Context().GetUserId(ctx),            
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sTitle)Delete(ctx context.Context, titleIds []int64) (err error){
    err = g.Try(ctx,func(ctx context.Context){
		var ids []int64
		ids = append(ids, titleIds...)        
        _, err = dao.Title.Ctx(ctx).Delete(dao.Title.Columns().TitleId+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}


