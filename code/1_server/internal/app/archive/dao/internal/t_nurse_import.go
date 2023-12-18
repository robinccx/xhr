// ==========================================================================
// GFast自动生成dao internal操作代码。
// 生成日期：2023-12-16 10:46:43
// 生成路径: internal/app/archive/dao/internal/t_nurse_import.go
// 生成人：xiao
// desc:护士导入
// company:云南奇讯科技有限公司
// ==========================================================================


package internal


import (
    "context"
    "github.com/gogf/gf/v2/database/gdb"
    "github.com/gogf/gf/v2/frame/g"
)


// NurseImportDao is the manager for logic model data accessing and custom defined data operations functions management.
type NurseImportDao struct {
    table   string         // Table is the underlying table name of the DAO.
    group   string         // Group is the database configuration group name of current DAO.
    columns NurseImportColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}


// NurseImportColumns defines and stores column names for table t_nurse_import.
type NurseImportColumns struct {    
    Id  string  //    
    SessionId  string  //    
    NurseId  string  //    
    NurseCode  string  //    
    NurseName  string  //    
    StartDate  string  //    
    EndDate  string  //    
    Sex  string  //    
    Birthday  string  //    
    AreaId  string  //    
    AreaCode  string  //    
    DeptId  string  //    
    DeptCode  string  //    
    EducationId  string  //    
    EducationCode  string  //    
    TitleId  string  //    
    TitleCode  string  //    
    StaffType  string  //    
    StaffTypeCode  string  //    
    CertNo  string  //    
    CertEnddate  string  //    
    RegDate  string  //    
    Note  string  //    
    Seq  string  //    
    Status  string  //    
    CreatedAt  string  //    
    CreatedBy  string  //    
}


var nurseimportColumns = NurseImportColumns{    
    Id:  "id",    
    SessionId:  "session_id",    
    NurseId:  "nurse_id",    
    NurseCode:  "nurse_code",    
    NurseName:  "nurse_name",    
    StartDate:  "start_date",    
    EndDate:  "end_date",    
    Sex:  "sex",    
    Birthday:  "birthday",    
    AreaId:  "area_id",    
    AreaCode:  "area_code",    
    DeptId:  "dept_id",    
    DeptCode:  "dept_code",    
    EducationId:  "education_id",    
    EducationCode:  "education_code",    
    TitleId:  "title_id",    
    TitleCode:  "title_code",    
    StaffType:  "staff_type",    
    StaffTypeCode:  "staff_type_code",    
    CertNo:  "cert_no",    
    CertEnddate:  "cert_end_date",    
    RegDate:  "reg_date",    
    Note:  "note",    
    Seq:  "seq",    
    Status:  "status",    
    CreatedAt:  "created_at",    
    CreatedBy:  "created_by",    
}


// NewNurseImportDao creates and returns a new DAO object for table data access.
func NewNurseImportDao() *NurseImportDao {
	return &NurseImportDao{
        group:    "default",
        table: "t_nurse_import",
        columns:nurseimportColumns,
	}
}


// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NurseImportDao) DB() gdb.DB {
	return g.DB(dao.group)
}


// Table returns the table name of current dao.
func (dao *NurseImportDao) Table() string {
	return dao.table
}


// Columns returns all column names of current dao.
func (dao *NurseImportDao) Columns() NurseImportColumns {
	return dao.columns
}


// Group returns the configuration group name of database of current dao.
func (dao *NurseImportDao) Group() string {
	return dao.group
}


// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NurseImportDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}


// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NurseImportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}