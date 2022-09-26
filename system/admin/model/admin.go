package model

import (
	"time"
)

//管理员表
type Admin struct {
    Id            int       `gorm:"column:id;primary_key" json:"id"`
    ParentId      int       `gorm:"column:parent_id" json:"parent_id"` //父级id
    HeadImg       string    `gorm:"column:head_img" json:"head_img"`   //用户头像
    Name          string    `gorm:"column:name" json:"name"`
    NickName      string    `gorm:"column:nick_name" json:"nick_name"` //昵称
    Gender        int8      `gorm:"column:gender" json:"gender"`       //性别 0=女 1=男 2=保密
    Password      string    `gorm:"-" json:"password"`
    Phone         string    `gorm:"column:phone" json:"phone"`                   //手机
    Email         string    `gorm:"column:email" json:"email"`                   //邮箱
    Status        int8      `gorm:"column:status" json:"status"`                 //状态 0=开启 1=关闭
    IsDelete      int8      `gorm:"column:is_delete" json:"is_delete"`           //是否删 0=否 1=是
    PromotionCode string    `gorm:"column:promotion_code" json:"promotion_code"` //推广码
    Info          string    `gorm:"column:info" json:"info"`                     //备注
    CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
    UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
}

var AdminGenderName = map[int8]string{
    0: "女",
    1: "男",
    2: "保密",
}
