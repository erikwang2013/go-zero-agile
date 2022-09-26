package model

// 管理员角色组
type AdminRoleGroup struct {
	Id int `gorm:"column:id;primary_key;type:int(11);not null" json:"id"`
	AdminId int `gorm:"column:admin_id" json:"admin_id"`
	RoleId int `gorm:"column:role_id" json:"role_id"`
	Status int8 `gorm:"column:status" json:"status"` //状态 0=开启 1=关闭
	IsDelete int8 `gorm:"column:is_delete" json:"is_delete"` //是否删 0=否 1=是
}