package model

import "time"

//登录日志
type AdminLoginLog struct {
    Id          uint64 `gorm:"column:id;primary_key" json:"id"`
    AdminId     int    `gorm:"column:admin_id" json:"admin_id"`
    AccessToken string `gorm:"column:access_token" json:"access_token"`
    LoginIp     string `gorm:"column:login_ip" json:"login_ip"`     //登录ip
    LoginTime   time.Time  `gorm:"column:login_time" json:"login_time"` //登录时间
}
