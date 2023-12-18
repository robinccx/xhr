// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/tiger1103/gfast/v3/api/v1/system"
	"github.com/tiger1103/gfast/v3/internal/app/system/model/entity"
)

type (
	ISysRole interface {
		GetRoleListSearch(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
		GetRoleList(ctx context.Context) (list []*entity.SysRole, err error)
		AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error)
		DelRoleRule(ctx context.Context, roleId int64) (err error)
		AddRole(ctx context.Context, req *system.RoleAddReq) (err error)
		Get(ctx context.Context, id uint) (res *entity.SysRole, err error)
		GetFilteredNamedPolicy(ctx context.Context, id uint) (gpSlice []int, err error)
		EditRole(ctx context.Context, req *system.RoleEditReq) (err error)
		DeleteByIds(ctx context.Context, ids []int64) (err error)
		RoleDeptTreeSelect(ctx context.Context, roleId int64) (res *system.RoleDeptTreeSelectRes, err error)
		GetRoleDepts(ctx context.Context, roleId int64) ([]int64, error)
		RoleDataScope(ctx context.Context, req *system.DataScopeReq) error
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}