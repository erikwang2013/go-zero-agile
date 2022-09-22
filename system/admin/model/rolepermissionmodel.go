package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RolePermissionModel = (*customRolePermissionModel)(nil)

type (
	// RolePermissionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRolePermissionModel.
	RolePermissionModel interface {
		rolePermissionModel
	}

	customRolePermissionModel struct {
		*defaultRolePermissionModel
	}
)

// NewRolePermissionModel returns a model for the database table.
func NewRolePermissionModel(conn sqlx.SqlConn, c cache.CacheConf) RolePermissionModel {
	return &customRolePermissionModel{
		defaultRolePermissionModel: newRolePermissionModel(conn, c),
	}
}
