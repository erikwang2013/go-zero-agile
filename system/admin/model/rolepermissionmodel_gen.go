// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rolePermissionFieldNames          = builder.RawFieldNames(&RolePermission{})
	rolePermissionRows                = strings.Join(rolePermissionFieldNames, ",")
	rolePermissionRowsExpectAutoSet   = strings.Join(stringx.Remove(rolePermissionFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	rolePermissionRowsWithPlaceHolder = strings.Join(stringx.Remove(rolePermissionFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheRolePermissionIdPrefix = "cache:rolePermission:id:"
)

type (
	rolePermissionModel interface {
		Insert(ctx context.Context, data *RolePermission) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RolePermission, error)
		Update(ctx context.Context, data *RolePermission) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRolePermissionModel struct {
		sqlc.CachedConn
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

func newRolePermissionModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultRolePermissionModel {
	return &defaultRolePermissionModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`role_permission`",
	}
}

func (m *defaultRolePermissionModel) Delete(ctx context.Context, id int64) error {
	rolePermissionIdKey := fmt.Sprintf("%s%v", cacheRolePermissionIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, rolePermissionIdKey)
	return err
}

func (m *defaultRolePermissionModel) FindOne(ctx context.Context, id int64) (*RolePermission, error) {
	rolePermissionIdKey := fmt.Sprintf("%s%v", cacheRolePermissionIdPrefix, id)
	var resp RolePermission
	err := m.QueryRowCtx(ctx, &resp, rolePermissionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolePermissionRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
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
	rolePermissionIdKey := fmt.Sprintf("%s%v", cacheRolePermissionIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, rolePermissionRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.RoleId, data.PermissionId, data.Status, data.IsDelete)
	}, rolePermissionIdKey)
	return ret, err
}

func (m *defaultRolePermissionModel) Update(ctx context.Context, data *RolePermission) error {
	rolePermissionIdKey := fmt.Sprintf("%s%v", cacheRolePermissionIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rolePermissionRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RoleId, data.PermissionId, data.Status, data.IsDelete, data.Id)
	}, rolePermissionIdKey)
	return err
}

func (m *defaultRolePermissionModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRolePermissionIdPrefix, primary)
}

func (m *defaultRolePermissionModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rolePermissionRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRolePermissionModel) tableName() string {
	return m.table
}
