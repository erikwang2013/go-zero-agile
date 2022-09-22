package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminRoleGroupModel = (*customAdminRoleGroupModel)(nil)

type (
	// AdminRoleGroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminRoleGroupModel.
	AdminRoleGroupModel interface {
		adminRoleGroupModel
	}

	customAdminRoleGroupModel struct {
		*defaultAdminRoleGroupModel
	}
)

// NewAdminRoleGroupModel returns a model for the database table.
func NewAdminRoleGroupModel(conn sqlx.SqlConn, c cache.CacheConf) AdminRoleGroupModel {
	return &customAdminRoleGroupModel{
		defaultAdminRoleGroupModel: newAdminRoleGroupModel(conn, c),
	}
}
