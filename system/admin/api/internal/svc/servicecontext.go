package svc

import (
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/middleware"
	"erik-agile/system/admin/api/internal/svc/gorm"

	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
    Config      config.Config
    AdminMiddle rest.Middleware
    // AdminModel          admin.AdminModel
    // AdminLoginLogModel  adminLoginLog.AdminLoginLogModel
    // AdminRoleGroupModel adminRoleGroup.AdminRoleGroupModel
    // RoleModel           role.RoleModel
    // PermissionModel     permission.PermissionModel
    // RolePermissionModel rolePermission.RolePermissionModel
}

// func NewServiceContext(c config.Config) *ServiceContext {
//     conn := sqlx.NewMysql(c.Mysql.DataSource)
//     return &ServiceContext{
//         Config:             c,
//         AdminMiddle:        middleware.NewAdminMiddleware().Handle,
//         AdminModel:         admin.NewAdminModel(conn, c.CacheRedis),
//         AdminLoginLogModel: adminLoginLog.NewAdminLoginLogModel(conn),
//         Gorm:               xgorm.NewGorm(c.Mysql.DataSource),
//     }
// }
func NewServiceContext(c config.Config, db *gorm.Gormdb) *ServiceContext {
    return &ServiceContext{
        Config:      c,
        AdminMiddle: middleware.NewAdminMiddleware(db).Handle,
    }
}
