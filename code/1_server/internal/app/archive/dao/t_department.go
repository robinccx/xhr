// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-12-08 09:19:50
// 生成路径: internal/app/archive/dao/t_department.go
// 生成人：xiao
// desc:科室
// company:云南奇讯科技有限公司
// ==========================================================================


package dao


import (
    "github.com/tiger1103/gfast/v3/internal/app/archive/dao/internal"
)
// departmentDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type departmentDao struct {
	*internal.DepartmentDao
}
var (
    // Department is globally public accessible object for table tools_gen_table operations.
    Department = departmentDao{
        internal.NewDepartmentDao(),
    }
)


// Fill with you ideas below.


