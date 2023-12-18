// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-11-30 18:32:51
// 生成路径: internal/app/archive/dao/internal/t_title.go
// 生成人：xiao
// desc:职称
// company:云南奇讯科技有限公司
// ==========================================================================


package internal


import (
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)


// TitleDao is the manager for logic model data accessing and custom defined data operations functions management.
type TitleDao struct {
    table   string         // Table is the underlying table name of the DAO.
    group   string         // Group is the database configuration group name of current DAO.
    columns TitleColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}


// TitleColumns defines and stores column names for table t_title.
type TitleColumns struct {    
    TitleId  string  // ID    
    TitleName  string  // 职称    
    Seq  string  // 次序    
    Status  string  // 状态    
    CreatedAt  string  // 创建时间    
    CreatedBy  string  // 创建用户    
    UpdatedAt  string  //    
    UpdatedBy  string  // 修改用户    
    DeletedAt  string  //    
    DeletedBy  string  // 删除用户    
}


var titleColumns = TitleColumns{    
    TitleId:  "title_id",    
    TitleName:  "title_name",    
    Seq:  "seq",    
    Status:  "status",    
    CreatedAt:  "created_at",    
    CreatedBy:  "created_by",    
    UpdatedAt:  "updated_at",    
    UpdatedBy:  "updated_by",    
    DeletedAt:  "deleted_at",    
    DeletedBy:  "deleted_by",    
}


// NewTitleDao creates and returns a new DAO object for table data access.
func NewTitleDao() *TitleDao {
	return &TitleDao{
        group:    "default",
        table: "t_title",
        columns:titleColumns,
	}
}


// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TitleDao) DB() gdb.DB {
	return g.DB(dao.group)
}


// Table returns the table name of current dao.
func (dao *TitleDao) Table() string {
	return dao.table
}


// Columns returns all column names of current dao.
func (dao *TitleDao) Columns() TitleColumns {
	return dao.columns
}


// Group returns the configuration group name of database of current dao.
func (dao *TitleDao) Group() string {
	return dao.group
}


// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TitleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}


// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TitleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}