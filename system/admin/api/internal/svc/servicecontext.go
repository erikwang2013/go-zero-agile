package svc

import (
	"go-zero-agile/system/admin/api/internal/config"
	"go-zero-agile/system/admin/api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
    AdminMiddle rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
        AdminMiddle:middleware.NewAdminMiddleware().Handle,
	}
}
