// Code generated by goctl. DO NOT EDIT!

package permission

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rolePermissionFieldNames          = builder.RawFieldNames(&RolePermission{})
	rolePermissionRows                = strings.Join(rolePermissionFieldNames, ",")
	rolePermissionRowsExpectAutoSet   = strings.Join(stringx.Remove(rolePermissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	rolePermissionRowsWithPlaceHolder = strings.Join(stringx.Remove(rolePermissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	rolePermissionModel interface {
		Insert(ctx context.Context, data *RolePermission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RolePermission, error)
		Update(ctx context.Context, data *RolePermission) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRolePermissionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RolePermission struct {
		Id           int64 `db:"id"`
		RoleId       int64 `db:"role_id"`
		PermissionId int64 `db:"permission_id"`
		Status       int64 `db:"status"`    // 状态 0=开启 1=关闭
		IsDelete     int64 `db:"is_delete"` // 是否删 0=否 1=是
	}
)

func newRolePermissionModel(conn sqlx.SqlConn) *defaultRolePermissionModel {
	return &defaultRolePermissionModel{
		conn:  conn,
		table: "`role_permission`",
	}
}

func (m *defaultRolePermissionModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRolePermissionModel) FindOne(ctx context.Context, id int64) (*RolePermission, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolePermissionRows, m.table)
	var resp RolePermission
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRolePermissionModel) Insert(ctx context.Context, data *RolePermission) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, rolePermissionRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.PermissionId, data.Status, data.IsDelete)
	return ret, err
}

func (m *defaultRolePermissionModel) Update(ctx context.Context, data *RolePermission) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rolePermissionRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.RoleId, data.PermissionId, data.Status, data.IsDelete, data.Id)
	return err
}

func (m *defaultRolePermissionModel) tableName() string {
	return m.table
}
