package model

import "time"

// 权限表
type Permission struct {
    Id               int       `gorm:"column:id"`
    ParentId         int       `gorm:"column:parent_id"` //父级
    Name             string    `gorm:"column:name"`      //权限名称
    ApiUrl           string    `gorm:"column:api_url"`   //api地址
    Code             string    `gorm:"column:code"`
    PermissionButton string    `gorm:"column:permission_button"` //权限按钮
    PermissionData   string    `gorm:"column:permission_data"`   //权限数据
    Info             string    `gorm:"column:info"`
    Status           int8      `gorm:"column:status"`    //状态 0=开启 1=关闭
    IsDelete         int8      `gorm:"column:is_delete"` //是否删 0=否 1=是
    CreateTime       time.Time `gorm:"column:create_time"`
}
