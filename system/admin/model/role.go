package model

import "time"

// 角色表
type Role struct {
    Id         int       `gorm:"column:id"`
    ParentId   int       `gorm:"column:parent_id"` //父级id   默认0为顶级
    Name       string    `gorm:"column:name"`
    Info       string    `gorm:"column:info"`
    Code       string    `gorm:"column:code"`
    Status     int8      `gorm:"column:status"`    //状态 0=开启 1=关闭
    IsDelete   int8      `gorm:"column:is_delete"` //是否删 0=否 1=是
    CreateTime time.Time `gorm:"column:create_time"`
}
