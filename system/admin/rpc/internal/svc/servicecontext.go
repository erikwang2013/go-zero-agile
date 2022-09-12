package svc

import (
	"go-zero-agile/system/admin/model"
	"go-zero-agile/system/admin/rpc/internal/config"

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
