// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-11-30 18:32:48
// 生成路径: internal/app/archive/logic/t_education.go
// 生成人：xiao
// desc:学历
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
	service.RegisterEducation(New())
}


func New() *sEducation {
	return &sEducation{}
}


type sEducation struct{}


func (s *sEducation)List(ctx context.Context, req *archive.EducationSearchReq) (listRes *archive.EducationSearchRes, err error){
    listRes = new(archive.EducationSearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.Education.Ctx(ctx).WithAll()        
        if req.EducationName != "" {
            m = m.Where(dao.Education.Columns().EducationName+" like ?", "%"+req.EducationName+"%")
        }          
        if req.Status != "" {
            m = m.Where(dao.Education.Columns().Status+" = ?", gconv.Int(req.Status))
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
        order:= "seq asc"
        if req.OrderBy!=""{
            order = req.OrderBy
        }
        var res []*model.EducationInfoRes
        err = m.Fields(archive.EducationSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)        
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.EducationListRes,len(res))
        for k,v:=range res{            
            listRes.List[k] = &model.EducationListRes{                  
                EducationId : v.EducationId,                    
                Seq : v.Seq,                    
                EducationName : v.EducationName,                    
                Status : v.Status,                
            }
        }
    })
    return
}


func (s *sEducation)GetByEducationId(ctx context.Context, educationId int64) (res *model.EducationInfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.Education.Ctx(ctx).WithAll().Where(dao.Education.Columns().EducationId,  educationId).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sEducation)Add(ctx context.Context, req *archive.EducationAddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {        
        _, err = dao.Education.Ctx(ctx).Insert(do.Education{            
            EducationName:req.EducationName,            
            Seq:req.Seq,            
            Status:req.Status,            
            CreatedBy:systemService.Context().GetUserId(ctx),            
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sEducation)Edit(ctx context.Context, req *archive.EducationEditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {        
        _, err = dao.Education.Ctx(ctx).WherePri(req.EducationId).Update(do.Education{            
            EducationName:req.EducationName,            
            Seq:req.Seq,            
            Status:req.Status,            
            UpdatedBy:systemService.Context().GetUserId(ctx),            
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sEducation)Delete(ctx context.Context, educationIds []int64) (err error){
    err = g.Try(ctx,func(ctx context.Context){
		var ids []int64
		ids = append(ids, educationIds...)        
        _, err = dao.Education.Ctx(ctx).Delete(dao.Education.Columns().EducationId+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}


