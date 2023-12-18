// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-11-30 18:32:51
// 生成路径: internal/app/archive/dao/t_title.go
// 生成人：xiao
// desc:职称
// company:云南奇讯科技有限公司
// ==========================================================================


package dao


import (
    "github.com/tiger1103/gfast/v3/internal/app/archive/dao/internal"
)
// titleDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type titleDao struct {
	*internal.TitleDao
}
var (
    // Title is globally public accessible object for table tools_gen_table operations.
    Title = titleDao{
        internal.NewTitleDao(),
    }
)


// Fill with you ideas below.


