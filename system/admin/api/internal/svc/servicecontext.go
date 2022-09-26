package svc

import (
	"erik-agile/common/xgorm"
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
    Config      config.Config
    AdminMiddle rest.Middleware
    Gorm        *gorm.DB
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
func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:      c,
        AdminMiddle: middleware.NewAdminMiddleware().Handle,
        Gorm:        xgorm.NewGorm(c.Mysql.DataSource, c.Mysql.TablePrefix),
    }
}
