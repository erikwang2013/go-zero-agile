// Code generated by goctl. DO NOT EDIT!

package admin

import (
    "context"
    "database/sql"
    "erik-agile/common/data-format"
    "fmt"
    "github.com/zeromicro/go-zero/core/logx"
    "github.com/zeromicro/go-zero/core/stores/builder"
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlc"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
    "github.com/zeromicro/go-zero/core/stringx"
    "strings"
    "time"
)

var (
    adminFieldNames          = builder.RawFieldNames(&Admin{})
    adminRows                = strings.Join(adminFieldNames, ",")
    adminRowsExpectAutoSet   = strings.Join(stringx.Remove(adminFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
    adminRowsWithPlaceHolder = strings.Join(stringx.Remove(adminFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

    cacheAdminIdPrefix   = "cache:admin:id:"
    cacheAdminNamePrefix = "cache:admin:name:"
    AdminGenderName      = map[int8]string{
        0: "女",
        1: "男",
        2: "保密",
    }
)

type (
    adminModel interface {
        Insert(ctx context.Context, data *Admin) (sql.Result, error)
        FindOne(ctx context.Context, id int) (*Admin, error)
        FindOneName(ctx context.Context, name string) (*Admin, error)
        All(ctx context.Context, data *Admin, page, limit int) ([]*Admin, error)
        Update(ctx context.Context, data *Admin) error
        Delete(ctx context.Context, id int) error
    }

    defaultAdminModel struct {
        sqlc.CachedConn
        table string
    }

    Admin struct {
        Id            int       `db:"id"`
        ParentId      int       `db:"parent_id"` // 父级id
        HeadImg       string    `db:"head_img"`  // 用户头像
        Name          string    `db:"name"`
        NickName      string    `db:"nick_name"` // 昵称
        Gender        int8      `db:"gender"`    // 性别 0=女 1=男 2=保密
        Password      string    `db:"password"`
        Phone         string    `db:"phone"`          // 手机
        Email         string    `db:"email"`          // 邮箱
        Status        int8      `db:"status"`         // 状态 0=开启 1=关闭
        IsDelete      int8      `db:"is_delete"`      // 是否删 0=否 1=是
        PromotionCode string    `db:"promotion_code"` // 推广码
        Info          string    `db:"info"`           // 备注
        CreateTime    time.Time `db:"create_time"`
        UpdateTime    time.Time `db:"update_time"`
    }

    setId struct {
        Id int `json:"id"`
    }
)

func newAdminModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAdminModel {
    return &defaultAdminModel{
        CachedConn: sqlc.NewConn(conn, c),
        table:      "`admin`",
    }
}

func (m *defaultAdminModel) Delete(ctx context.Context, id int) error {
    adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
        query := fmt.Sprintf("delete from %s where `id` = ? and is_delete=1 ", m.table)
        return conn.ExecCtx(ctx, query, id)
    }, adminIdKey)
    return err
}

func (m *defaultAdminModel) All(ctx context.Context, data *Admin, page, limit int) ([]*Admin, error) {
    sql := "select %s from %s where is_delete=0"
    str := []string{}
    if len(data.NickName) > 0 {
        sql += " like nick_name ?"
        str = append(str, data.NickName)
    }
    if len(data.Name) > 0 {
        sql += " like name ?"
        str = append(str, data.Name)
    }
    if len(data.Phone) > 0 {
        sql += " like phone ?"
        str = append(str, data.Phone)
    }
    if len(data.Email) > 0 {
        sql += " like email ?"
        str = append(str, data.Email)
    }
    if data.Id > 0 {
        sql += " and id=?"
        str = append(str, dataFormat.IntToString(data.Id))
    }
    if data.ParentId >= 0 {
        sql += " and parent_id=?"
        str = append(str, dataFormat.IntToString(data.ParentId))
    }
    if data.Gender >= 0 {
        sql += " and gender=?"
        str = append(str, dataFormat.IntToString(int(data.Gender)))
    }
    if data.Status >= 0 {
        sql += " and status=?"
        str = append(str, dataFormat.IntToString(int(data.Status)))
    }
    var resp []*Admin
    query := fmt.Sprintf(sql, adminRows, m.table)
    logx.Info("===查询sql==")
    logx.Info(query)
    err := m.QueryRowsNoCache(&resp, query, str)
    logx.Info("===查询报错==")
    logx.Info(err)
    switch err {
    case nil:
        return resp, nil
    case sqlc.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

func (m *defaultAdminModel) FindOneName(ctx context.Context, name string) (*Admin, error) {
    adminNameKey := fmt.Sprintf("%s%v", cacheAdminNamePrefix, name)
    var id setId
    err := m.GetCache(adminNameKey, &id)
    var resp Admin
    if err != nil {
        query := fmt.Sprintf("select %s from %s where `name` = ? and is_delete=0 limit 1", adminRows, m.table)
        err := m.QueryRowNoCache(&resp, query, name)
        switch err {
        case nil:
            adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, resp.Id)
            m.SetCache(adminIdKey, resp)
            id.Id = resp.Id
            m.SetCache(adminNameKey, id)
            return &resp, nil
        case sqlc.ErrNotFound:
            return nil, ErrNotFound
        default:
            return nil, err
        }
    } else {
        adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, id.Id)
        err := m.GetCache(adminIdKey, &resp)
        switch err {
        case nil:
            return &resp, nil
        case sqlc.ErrNotFound:
            return nil, ErrNotFound
        default:
            return nil, err
        }
    }

}

func (m *defaultAdminModel) FindOne(ctx context.Context, id int) (*Admin, error) {
    adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, id)
    var resp Admin
    err := m.QueryRowCtx(ctx, &resp, adminIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
        query := fmt.Sprintf("select %s from %s where `id` = ? and is_delete=0 limit 1", adminRows, m.table)
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

func (m *defaultAdminModel) Insert(ctx context.Context, data *Admin) (sql.Result, error) {
    adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, data.Id)
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
        query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, adminRowsExpectAutoSet)
        return conn.ExecCtx(ctx, query, data.Id, data.ParentId, data.HeadImg, data.Name, data.NickName, data.Gender, data.Password, data.Phone, data.Email, data.Status, data.IsDelete, data.PromotionCode, data.Info)
    }, adminIdKey)
    return ret, err
}

func (m *defaultAdminModel) Update(ctx context.Context, data *Admin) error {
    adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, data.Id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
        query := fmt.Sprintf("update %s set %s where `id` = ? and is_delete=0 ", m.table, adminRowsWithPlaceHolder)
        return conn.ExecCtx(ctx, query, data.ParentId, data.HeadImg, data.Name, data.NickName, data.Gender, data.Password, data.Phone, data.Email, data.Status, data.IsDelete, data.PromotionCode, data.Info, data.Id)
    }, adminIdKey)
    return err
}

func (m *defaultAdminModel) formatPrimary(primary interface{}) string {
    return fmt.Sprintf("%s%v", cacheAdminIdPrefix, primary)
}

func (m *defaultAdminModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
    query := fmt.Sprintf("select %s from %s where `id` = ? and is_delete=0 limit 1", adminRows, m.table)
    return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAdminModel) tableName() string {
    return m.table
}
