package log

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminLoginLogModel = (*customAdminLoginLogModel)(nil)

type (
	// AdminLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminLoginLogModel.
	AdminLoginLogModel interface {
		adminLoginLogModel
	}

	customAdminLoginLogModel struct {
		*defaultAdminLoginLogModel
	}
)

// NewAdminLoginLogModel returns a model for the database table.
func NewAdminLoginLogModel(conn sqlx.SqlConn) AdminLoginLogModel {
	return &customAdminLoginLogModel{
		defaultAdminLoginLogModel: newAdminLoginLogModel(conn),
	}
}
