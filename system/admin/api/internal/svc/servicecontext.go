package svc

import (
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/middleware"
	"erik-agile/system/admin/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
    Config      config.Config
    AdminMiddle rest.Middleware
    AdminModel  model.AdminModel
}

func NewServiceContext(c config.Config) *ServiceContext {
    conn := sqlx.NewMysql(c.Mysql.DataSource)
    return &ServiceContext{
        Config:      c,
        AdminMiddle: middleware.NewAdminMiddleware().Handle,
        AdminModel:  model.NewAdminModel(conn, c.CacheRedis),
    }
}
