// ==========================================================================
// GFast自动生成logic操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/logic/t_department.go
// 生成人：xiao
// desc:科室
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
	service.RegisterDepartment(New())
}


func New() *sDepartment {
	return &sDepartment{}
}


type sDepartment struct{}


func (s *sDepartment)List(ctx context.Context, req *archive.DepartmentSearchReq) (listRes *archive.DepartmentSearchRes, err error){
    listRes = new(archive.DepartmentSearchRes)
    err = g.Try(ctx, func(ctx context.Context) {
        m := dao.Department.Ctx(ctx).WithAll()          
        if req.DeptCode != "" {
            m = m.Where(dao.Department.Columns().DeptCode+" = ?", req.DeptCode)
        }        
        if req.DeptName != "" {
            m = m.Where(dao.Department.Columns().DeptName+" like ?", "%"+req.DeptName+"%")
        }          
        if req.Status != "" {
            m = m.Where(dao.Department.Columns().Status+" = ?", gconv.Int(req.Status))
        }          
        if req.DeptType != "" {
            m = m.Where(dao.Department.Columns().DeptType+" = ?", req.DeptType)
        }          
        if req.Att1 != "" {
            m = m.Where(dao.Department.Columns().Att1+" = ?", req.Att1)
        }          
        if req.Att2 != "" {
            m = m.Where(dao.Department.Columns().Att2+" = ?", req.Att2)
        }          
        if req.Att3 != "" {
            m = m.Where(dao.Department.Columns().Att3+" = ?", req.Att3)
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
        order:= "dept_code asc"
        if req.OrderBy!=""{
            order = req.OrderBy
        }
        var res []*model.DepartmentInfoRes
        err = m.Fields(archive.DepartmentSearchRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)        
        liberr.ErrIsNil(ctx, err, "获取数据失败")
        listRes.List = make([]*model.DepartmentListRes,len(res))
        for k,v:=range res{            
            listRes.List[k] = &model.DepartmentListRes{                  
                DeptId : v.DeptId,                    
                DeptCode : v.DeptCode,                    
                DeptName : v.DeptName,                    
                DeptType : v.DeptType,                    
                Leader : v.Leader,                    
                Phone : v.Phone,                    
                Email : v.Email,                    
                Att1 : v.Att1,                    
                Att2 : v.Att2,                    
                Att3 : v.Att3,                    
                Status : v.Status,                    
                CreatedAt : v.CreatedAt,                    
                CreatedBy : v.CreatedBy,                
            }
        }
    })
    return
}


func (s *sDepartment)GetByDeptId(ctx context.Context, deptId int64) (res *model.DepartmentInfoRes,err error){
    err =g.Try(ctx, func(ctx context.Context){
        err = dao.Department.Ctx(ctx).WithAll().Where(dao.Department.Columns().DeptId,  deptId).Scan(&res)
        liberr.ErrIsNil(ctx,err,"获取信息失败")
    })
    return
}


func (s *sDepartment)Add(ctx context.Context, req *archive.DepartmentAddReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {        
        _, err = dao.Department.Ctx(ctx).Insert(do.Department{            
            DeptCode:req.DeptCode,            
            DeptName:req.DeptName,            
            DeptType:req.DeptType,            
            Leader:req.Leader,            
            Phone:req.Phone,            
            Email:req.Email,            
            Att1:req.Att1,            
            Att2:req.Att2,            
            Att3:req.Att3,            
            Status:req.Status,            
            CreatedBy:systemService.Context().GetUserId(ctx),            
        })
        liberr.ErrIsNil(ctx, err, "添加失败")
    })
    return
}


func (s *sDepartment)Edit(ctx context.Context, req *archive.DepartmentEditReq) (err error){
    err = g.Try(ctx, func(ctx context.Context) {        
        _, err = dao.Department.Ctx(ctx).WherePri(req.DeptId).Update(do.Department{            
            DeptCode:req.DeptCode,            
            DeptName:req.DeptName,            
            DeptType:req.DeptType,            
            Leader:req.Leader,            
            Phone:req.Phone,            
            Email:req.Email,            
            Att1:req.Att1,            
            Att2:req.Att2,            
            Att3:req.Att3,            
            Status:req.Status,            
            UpdatedBy:systemService.Context().GetUserId(ctx),            
        })
        liberr.ErrIsNil(ctx, err, "修改失败")
    })
    return
}


func (s *sDepartment)Delete(ctx context.Context, deptIds []int64) (err error){
    err = g.Try(ctx,func(ctx context.Context){
		var ids []int64
		ids = append(ids, deptIds...)        
        _, err = dao.Department.Ctx(ctx).Delete(dao.Department.Columns().DeptId+" in (?)", ids)
        liberr.ErrIsNil(ctx,err,"删除失败")
    })
    return
}


