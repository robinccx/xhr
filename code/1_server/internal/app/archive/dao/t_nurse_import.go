// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-12-16 10:46:43
// 生成路径: internal/app/archive/dao/t_nurse_import.go
// 生成人：xiao
// desc:护士导入
// company:云南奇讯科技有限公司
// ==========================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/archive/dao/internal"
)

// nurseimportDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type nurseImportDao struct {
	*internal.NurseImportDao
}

var (
	// NurseImport is globally public accessible object for table tools_gen_table operations.
	NurseImport = nurseImportDao{
		internal.NewNurseImportDao(),
	}
)

// Fill with you ideas below.
