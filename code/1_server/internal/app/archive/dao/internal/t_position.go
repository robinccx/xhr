// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/dao/internal/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================


package internal


import (
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)


// PositionDao is the manager for logic model data accessing and custom defined data operations functions management.
type PositionDao struct {
    table   string         // Table is the underlying table name of the DAO.
    group   string         // Group is the database configuration group name of current DAO.
    columns PositionColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}


// PositionColumns defines and stores column names for table t_position.
type PositionColumns struct {    
    PositionId  string  // 职务ID    
    PositionCode  string  // 职务代码    
    PositionName  string  // 职务    
    Seq  string  // 次序    
    Status  string  // 状态    
    CreatedAt  string  // 创建时间    
    CreatedBy  string  // 创建用户    
    UpdatedAt  string  //    
    UpdatedBy  string  //    
    DeletedAt  string  //    
    DeletedBy  string  //    
}


var positionColumns = PositionColumns{    
    PositionId:  "position_id",    
    PositionCode:  "position_code",    
    PositionName:  "position_name",    
    Seq:  "seq",    
    Status:  "status",    
    CreatedAt:  "created_at",    
    CreatedBy:  "created_by",    
    UpdatedAt:  "updated_at",    
    UpdatedBy:  "updated_by",    
    DeletedAt:  "deleted_at",    
    DeletedBy:  "deleted_by",    
}


// NewPositionDao creates and returns a new DAO object for table data access.
func NewPositionDao() *PositionDao {
	return &PositionDao{
        group:    "default",
        table: "t_position",
        columns:positionColumns,
	}
}


// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PositionDao) DB() gdb.DB {
	return g.DB(dao.group)
}


// Table returns the table name of current dao.
func (dao *PositionDao) Table() string {
	return dao.table
}


// Columns returns all column names of current dao.
func (dao *PositionDao) Columns() PositionColumns {
	return dao.columns
}


// Group returns the configuration group name of database of current dao.
func (dao *PositionDao) Group() string {
	return dao.group
}


// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PositionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}


// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PositionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}