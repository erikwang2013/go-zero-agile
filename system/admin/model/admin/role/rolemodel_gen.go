// Code generated by goctl. DO NOT EDIT!

package role

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	roleFieldNames          = builder.RawFieldNames(&Role{})
	roleRows                = strings.Join(roleFieldNames, ",")
	roleRowsExpectAutoSet   = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	roleRowsWithPlaceHolder = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	roleModel interface {
		Insert(ctx context.Context, data *Role) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Role, error)
		Update(ctx context.Context, data *Role) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Role struct {
		Id         int64     `db:"id"`
		ParentId   int64     `db:"parent_id"` // 父级id   默认0为顶级
		Name       string    `db:"name"`
		Info       string    `db:"info"`
		Code       string    `db:"code"`
		Status     int64     `db:"status"`    // 状态 0=开启 1=关闭
		IsDelete   int64     `db:"is_delete"` // 是否删 0=否 1=是
		CreateTime time.Time `db:"create_time"`
	}
)

func newRoleModel(conn sqlx.SqlConn) *defaultRoleModel {
	return &defaultRoleModel{
		conn:  conn,
		table: "`role`",
	}
}

func (m *defaultRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id int64) (*Role, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
	var resp Role
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

func (m *defaultRoleModel) Insert(ctx context.Context, data *Role) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, roleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.Info, data.Code, data.Status, data.IsDelete)
	return ret, err
}

func (m *defaultRoleModel) Update(ctx context.Context, data *Role) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ParentId, data.Name, data.Info, data.Code, data.Status, data.IsDelete, data.Id)
	return err
}

func (m *defaultRoleModel) tableName() string {
	return m.table
}
