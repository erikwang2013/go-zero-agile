package svc

import (
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/middleware"
	"erik-agile/system/admin/model/admin"
	adminLoginLog "erik-agile/system/admin/model/admin/login/log"
	"erik-agile/system/admin/model/admin/permission"
	"erik-agile/system/admin/model/admin/role"
	adminRoleGroup "erik-agile/system/admin/model/admin/role/group"
	rolePermission "erik-agile/system/admin/model/admin/role/permission"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
    Config             config.Config
    AdminMiddle        rest.Middleware
    AdminModel         admin.AdminModel
    AdminLoginLogModel adminLoginLog.AdminLoginLogModel
    AdminRoleGroup     adminRoleGroup.AdminRoleGroupModel
    Role               role.RoleModel
    Permission         permission.PermissionModel
    RolePermission     rolePermission.RolePermissionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
    conn := sqlx.NewMysql(c.Mysql.DataSource)
    return &ServiceContext{
        Config:      c,
        AdminMiddle: middleware.NewAdminMiddleware().Handle,
        AdminModel:  admin.NewAdminModel(conn, c.CacheRedis),
    }
}
