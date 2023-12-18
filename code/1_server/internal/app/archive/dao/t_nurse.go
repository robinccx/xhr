// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-12-16 10:46:42
// 生成路径: internal/app/archive/dao/t_nurse.go
// 生成人：xiao
// desc:护士
// company:云南奇讯科技有限公司
// ==========================================================================


package dao


import (
    "github.com/tiger1103/gfast/v3/internal/app/archive/dao/internal"
)
// nurseDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type nurseDao struct {
	*internal.NurseDao
}
var (
    // Nurse is globally public accessible object for table tools_gen_table operations.
    Nurse = nurseDao{
        internal.NewNurseDao(),
    }
)


// Fill with you ideas below.


