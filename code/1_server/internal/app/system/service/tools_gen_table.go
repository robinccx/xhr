// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type (
	IToolsGenTable interface {
		List(ctx context.Context, req *system.ToolsGenTableSearchReq) (res *system.ToolsGenTableSearchRes, err error)
		SelectDbTableList(ctx context.Context, req *system.ToolsGenTableImportSearchReq) (res *system.ToolsGenTableSearchRes, err error)
		SelectDbTableListByNames(ctx context.Context, tableNames []string) ([]*entity.ToolsGenTable, error)
		ImportGenTable(ctx context.Context, tableList []*entity.ToolsGenTable) error
		InitTable(ctx context.Context, table *entity.ToolsGenTable, columns []*entity.ToolsGenTableColumn) error
		ConvertClassName(ctx context.Context, tableName string) string
		GetBusinessName(ctx context.Context, tableName string) string
		DeleteTable(ctx context.Context, req *system.ToolsGenTableDeleteReq) error
		ColumnList(ctx context.Context, req *system.ToolsGenTableEditReq) (res *system.ToolsGenTableEditRes, err error)
		GetTableInfoByTableId(ctx context.Context, tableId int64) (data *entity.ToolsGenTable, err error)
		GetRelationTable(ctx context.Context) (res []*model.ToolsGenTableColumnsData, err error)
		SaveEdit(ctx context.Context, req *system.ToolsGenTableColumnsEditReq) (err error)
		GenData(ctx context.Context, tableId int64) (data g.MapStrStr, extendData *model.ToolsGenTableEx, err error)
		SelectRecordById(ctx context.Context, tableId int64) (tableEx *model.ToolsGenTableEx, err error)
		GenCode(ctx context.Context, ids []int) (err error)
	}
)

var (
	localToolsGenTable IToolsGenTable
)

func ToolsGenTable() IToolsGenTable {
	if localToolsGenTable == nil {
		panic("implement not found for interface IToolsGenTable, forgot register?")
	}
	return localToolsGenTable
}

func RegisterToolsGenTable(i IToolsGenTable) {
	localToolsGenTable = i
}