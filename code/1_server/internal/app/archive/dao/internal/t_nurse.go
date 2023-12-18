// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-12-16 10:46:42
// 生成路径: internal/app/archive/dao/internal/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================


package internal


import (
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)


// NurseDao is the manager for logic model data accessing and custom defined data operations functions management.
type NurseDao struct {
    table   string         // Table is the underlying table name of the DAO.
    group   string         // Group is the database configuration group name of current DAO.
    columns NurseColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}


// NurseColumns defines and stores column names for table t_nurse.
type NurseColumns struct {    
    NurseId  string  // ID    
    NurseCode  string  // 编号    
    NurseName  string  // 名称    
    Sex  string  // 性别    
    Birthday  string  // 出生年月    
    AreaId  string  // 院区    
    DeptId  string  // 科室    
    StartDate  string  // 入职日期    
    EducationId  string  // 学历    
    TitleId  string  // 职称    
    StaffType  string  // 人员类别    
    Note  string  // 备注    
    EndDate  string  // 离职日期    
    Status  string  // 状态    
    CreatedAt  string  // 创建时间    
    CreatedBy  string  // 创建用户    
    UpdatedAt  string  //    
    UpdatedBy  string  // 修改用户    
    DeletedAt  string  //    
    DeletedBy  string  // 删除用户    
    SessionId  string  //    
    CertEnddate  string  // 职业证书到期日期    
    RegDate  string  // 注册或变更日期    
    CertNo  string  // 职业证书    
    IdNo  string  // 身份证号    
    PositionId  string  // 职务    
}


var nurseColumns = NurseColumns{    
    NurseId:  "nurse_id",    
    NurseCode:  "nurse_code",    
    NurseName:  "nurse_name",    
    Sex:  "sex",    
    Birthday:  "birthday",    
    AreaId:  "area_id",    
    DeptId:  "dept_id",    
    StartDate:  "start_date",    
    EducationId:  "education_id",    
    TitleId:  "title_id",    
    StaffType:  "staff_type",    
    Note:  "note",    
    EndDate:  "end_date",    
    Status:  "status",    
    CreatedAt:  "created_at",    
    CreatedBy:  "created_by",    
    UpdatedAt:  "updated_at",    
    UpdatedBy:  "updated_by",    
    DeletedAt:  "deleted_at",    
    DeletedBy:  "deleted_by",    
    SessionId:  "session_id",    
    CertEnddate:  "cert_end_date",    
    RegDate:  "reg_date",    
    CertNo:  "cert_no",    
    IdNo:  "id_no",    
    PositionId:  "position_id",    
}


// NewNurseDao creates and returns a new DAO object for table data access.
func NewNurseDao() *NurseDao {
	return &NurseDao{
        group:    "default",
        table: "t_nurse",
        columns:nurseColumns,
	}
}


// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NurseDao) DB() gdb.DB {
	return g.DB(dao.group)
}


// Table returns the table name of current dao.
func (dao *NurseDao) Table() string {
	return dao.table
}


// Columns returns all column names of current dao.
func (dao *NurseDao) Columns() NurseColumns {
	return dao.columns
}


// Group returns the configuration group name of database of current dao.
func (dao *NurseDao) Group() string {
	return dao.group
}


// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NurseDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}


// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NurseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}