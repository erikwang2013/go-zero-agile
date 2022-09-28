package model

// 角色权限关系表
type RolePermission struct {
    Id                  int    `gorm:"column:id"`
    RoleId              int    `gorm:"column:role_id"`
    PermissionId        int    `gorm:"column:permission_id"`
    PermissionChilderId string `gorm:"column:permission_childer_id"` //子权限
    PermissionButton    string `gorm:"column:permission_button"`     //按钮json
    Status              int8   `gorm:"column:status"`                //状态 0=开启 1=关闭
    IsDelete            int8   `gorm:"column:is_delete"`             //是否删 0=否 1=是
}
