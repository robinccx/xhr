// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/dao/internal/t_department.go
// 生成人：xiao
// desc:科室
// company:云南奇讯科技有限公司
// ==========================================================================


package internal


import (
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)


// DepartmentDao is the manager for logic model data accessing and custom defined data operations functions management.
type DepartmentDao struct {
    table   string         // Table is the underlying table name of the DAO.
    group   string         // Group is the database configuration group name of current DAO.
    columns DepartmentColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}


// DepartmentColumns defines and stores column names for table t_department.
type DepartmentColumns struct {    
    DeptId  string  // 科室ID    
    ParentId  string  // 父级ID    
    DeptCode  string  // 科室代码    
    DeptName  string  // 科室名称    
    DeptType  string  // 片区    
    Leader  string  // 负责人    
    Phone  string  // 电话    
    Email  string  // 邮件    
    Att1  string  // 属性1    
    Att2  string  // 属性2    
    Att3  string  // 属性3    
    Status  string  // 状态    
    CreatedAt  string  // 创建时间    
    CreatedBy  string  // 创建用户    
    UpdatedAt  string  //    
    UpdatedBy  string  // 修改用户    
    DeletedAt  string  //    
    DeletedBy  string  // 删除用户    
}


var departmentColumns = DepartmentColumns{    
    DeptId:  "dept_id",    
    ParentId:  "parent_id",    
    DeptCode:  "dept_code",    
    DeptName:  "dept_name",    
    DeptType:  "dept_type",    
    Leader:  "leader",    
    Phone:  "phone",    
    Email:  "email",    
    Att1:  "att1",    
    Att2:  "att2",    
    Att3:  "att3",    
    Status:  "status",    
    CreatedAt:  "created_at",    
    CreatedBy:  "created_by",    
    UpdatedAt:  "updated_at",    
    UpdatedBy:  "updated_by",    
    DeletedAt:  "deleted_at",    
    DeletedBy:  "deleted_by",    
}


// NewDepartmentDao creates and returns a new DAO object for table data access.
func NewDepartmentDao() *DepartmentDao {
	return &DepartmentDao{
        group:    "default",
        table: "t_department",
        columns:departmentColumns,
	}
}


// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DepartmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}


// Table returns the table name of current dao.
func (dao *DepartmentDao) Table() string {
	return dao.table
}


// Columns returns all column names of current dao.
func (dao *DepartmentDao) Columns() DepartmentColumns {
	return dao.columns
}


// Group returns the configuration group name of database of current dao.
func (dao *DepartmentDao) Group() string {
	return dao.group
}


// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DepartmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}


// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DepartmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}