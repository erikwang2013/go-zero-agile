package svc

import (
	"erik-agile/system/admin/model"
	"erik-agile/system/admin/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
    AdminModel model.AdminModel
}

func NewServiceContext(c config.Config) *ServiceContext {
    conn:=sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
        AdminModel:model.NewAdminModel(conn, c.CacheRedis),
	}
}
