package model

import "time"

//管理员表
type Admin struct {
    Id            int       `gorm:"column:id;primary_key"`
    ParentId      int       `gorm:"column:parent_id"` //父级id
    HeadImg       string    `gorm:"column:head_img"`  //用户头像
    Name          string    `gorm:"column:name"`
    NickName      string    `gorm:"column:nick_name"` //昵称
    Gender        int8      `gorm:"column:gender"`    //性别 0=女 1=男 2=保密
    Password      string    `gorm:"column:password;-"`
    Phone         string    `gorm:"column:phone"`          //手机
    Email         string    `gorm:"column:email"`          //邮箱
    Status        int8      `gorm:"column:status"`         //状态 0=开启 1=关闭
    IsDelete      int8      `gorm:"column:is_delete"`      //是否删 0=否 1=是
    PromotionCode string    `gorm:"column:promotion_code"` //推广码
    Info          string    `gorm:"column:info"`           //备注
    CreateTime    time.Time `gorm:"column:create_time"`
    UpdateTime    time.Time `gorm:"column:update_time"`
}

var AdminGenderName = map[int8]string{
    0: "女",
    1: "男",
    2: "保密",
}
