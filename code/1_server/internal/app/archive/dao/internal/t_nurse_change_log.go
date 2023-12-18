// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-12-11 10:46:35
// 生成路径: internal/app/archive/dao/internal/t_nurse_change_log.go
// 生成人：xiao
// desc:变更记录查询
// company:云南奇讯科技有限公司
// ==========================================================================


package internal


import (
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)


// NurseChangeLogDao is the manager for logic model data accessing and custom defined data operations functions management.
type NurseChangeLogDao struct {
    table   string         // Table is the underlying table name of the DAO.
    group   string         // Group is the database configuration group name of current DAO.
    columns NurseChangeLogColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}


// NurseChangeLogColumns defines and stores column names for table t_nurse_change_log.
type NurseChangeLogColumns struct {    
    Id  string  //    
    NurseId  string  //    
    ChangeType  string  //    
    FromAreaId  string  //    
    FromDeptId  string  //    
    FromEntityId  string  //    
    FromEntityName  string  //    
    ToAreaId  string  //    
    ToDeptId  string  //    
    ToEntityId  string  //    
    ToEntityName  string  //    
    TransDate  string  //    
    Status  string  //    
    CreatedAt  string  //    
    CreatedBy  string  //    
    UpdatedAt  string  //    
    UpdatedBy  string  //    
    DeletedAt  string  //    
    DeletedBy  string  //    
}


var nurseChangeLogColumns = NurseChangeLogColumns{    
    Id:  "id",    
    NurseId:  "nurse_id",    
    ChangeType:  "change_type",    
    FromAreaId:  "from_area_id",    
    FromDeptId:  "from_dept_id",    
    FromEntityId:  "from_entity_id",    
    FromEntityName:  "from_entity_name",    
    ToAreaId:  "to_area_id",    
    ToDeptId:  "to_dept_id",    
    ToEntityId:  "to_entity_id",    
    ToEntityName:  "to_entity_name",    
    TransDate:  "trans_date",    
    Status:  "status",    
    CreatedAt:  "created_at",    
    CreatedBy:  "created_by",    
    UpdatedAt:  "updated_at",    
    UpdatedBy:  "updated_by",    
    DeletedAt:  "deleted_at",    
    DeletedBy:  "deleted_by",    
}


// NewNurseChangeLogDao creates and returns a new DAO object for table data access.
func NewNurseChangeLogDao() *NurseChangeLogDao {
	return &NurseChangeLogDao{
        group:    "default",
        table: "t_nurse_change_log",
        columns:nurseChangeLogColumns,
	}
}


// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NurseChangeLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}


// Table returns the table name of current dao.
func (dao *NurseChangeLogDao) Table() string {
	return dao.table
}


// Columns returns all column names of current dao.
func (dao *NurseChangeLogDao) Columns() NurseChangeLogColumns {
	return dao.columns
}


// Group returns the configuration group name of database of current dao.
func (dao *NurseChangeLogDao) Group() string {
	return dao.group
}


// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NurseChangeLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}


// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NurseChangeLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}