// ==========================================================================
// GFast自动生成dao操作代码。
// 生成日期：2023-12-16 10:37:07
// 生成路径: internal/app/archive/dao/t_position.go
// 生成人：xiao
// desc:职务
// company:云南奇讯科技有限公司
// ==========================================================================


package dao


import (
    "github.com/tiger1103/gfast/v3/internal/app/archive/dao/internal"
)
// positionDao is the manager for logic model data accessing and custom defined data operations functions management.
// You can define custom methods on it to extend its functionality as you wish.
type positionDao struct {
	*internal.PositionDao
}
var (
    // Position is globally public accessible object for table tools_gen_table operations.
    Position = positionDao{
        internal.NewPositionDao(),
    }
)


// Fill with you ideas below.


