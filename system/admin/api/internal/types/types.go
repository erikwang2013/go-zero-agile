// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
    UserName string `json:"user_name" validate:"required,alphanum,max=20,min=4"`
    Password string `json:"password" validate:"required,alphanum,max=30,min=6"`
}

type LoginReply struct {
    Id           int    `json:"id"`
    Name         string `json:"name"`
    AccessToken  string `json:"access_token"`
    AccessExpire int64  `json:"access_expire"`
    RefreshAfter int64  `json:"refresh_after"`
}

type AdminInfoReq struct {
    Id       int    `json:"id,optional" validate:"gte=0"`
    ParentId int    `json:"parent_id,optional" validate:"number,max=18,min=0"`
    NickName string `json:"nick_name,optional" validate:"max=30,min=4"`
    Name     string `json:"name,optional" validate:"alphanum,max=30,min=4"`
    Phone    string `json:"phone,optional" validate:"number,len=11"` // 手机
    Email    string `json:"email,optional" validate:"email"`         // 邮箱
    Status   int8   `json:"status,optional" validate:"number,min=0,max=1"`
    Gender   int8   `json:"gender,optional" validate:"gte=0,lte=2"`
    Page     int    `json:"page" validate:"number,max=11,min=1"`
    Limit    int    `json:"limit" validate:"number,max=11,min=1"`
}

type AdminPutReq struct {
    Id       int    `json:"id" validate:"required,gte=0"`
    ParentId int    `json:"parent_id,optional" validate:"number,max=18,min=0"`
    NickName string `json:"nick_name,optional" validate:"max=30,min=4"`
    Name     string `json:"name,optional" validate:"alphanum,max=30,min=4"`
    Password string `json:"password,optional" validate:"alphanum,max=30,min=6"`
    Phone    string `json:"phone,optional" validate:"number,len=11"` // 手机
    Email    string `json:"email,optional" validate:"email"`         // 邮箱
    Status   int8   `json:"status,optional" validate:"number,min=0,max=1"`
    Gender   int8   `json:"gender,optional" validate:"gte=0,lte=2"`
    Info     string `json:"info,optional" validate:"max=100"` // 备注
}

type AdminDeleteReq struct {
    Id string `json:"id,optional"`
}
type AdminInfoReply struct {
    Id            int             `json:"id"`
    ParentId      int             `json:"parent_id"` // 父级id
    HeadImg       string          `json:"head_img"`  // 用户头像
    Name          string          `json:"name"`
    NickName      string          `json:"nick_name"`      // 昵称
    Gender        StatusValueName `json:"gender"`         // 性别 0=女 1=男 2=保密
    Phone         string          `json:"phone"`          // 手机
    Email         string          `json:"email"`          // 邮箱
    Status        StatusValueName `json:"status"`         // 状态 0=开启 1=关闭
    IsDelete      StatusValueName `json:"is_delete"`      // 是否删 0=否 1=是
    PromotionCode string          `json:"promotion_code"` // 推广码
    Info          string          `json:"info"`           // 备注
    CreateTime    int64           `json:"create_time"`
    UpdateTime    int64           `json:"update_time"`
}

type StatusValueName struct {
    Key int8   `json:"key"`
    Val string `json:"val"`
}

type AdminAddReq struct {
    ParentId int    `json:"parent_id" validate:"required,gte=0"`    // 父级id
    HeadImg  string `json:"head_img" validate:"url,max=250,min=10"` // 用户头像
    Name     string `json:"name" validate:"required,alphanum,max=30,min=4"`
    Password string `json:"password" validate:"required,alphanum,max=30,min=6"`
    NickName string `json:"nick_name" validate:"max=30,min=4"`       // 昵称
    Gender   int8   `json:"gender" validate:"gte=0,lte=2"`           // 性别 0=女 1=男 2=保密
    Phone    string `json:"phone" validate:"required,number,len=11"` // 手机
    Email    string `json:"email" validate:"required,email"`         // 邮箱
    Status   int8   `json:"status" validate:"number,min=0,max=1"`    // 状态 0=开启 1=关闭
    Info     string `json:"info" validate:"max=100"`                 // 备注
}

type AdminInfoAllReq struct {
    Id int `json:"id" validate:"required,gte=0"`
}

type PermissionAddReq struct {
    ParentId         int              `json:"parent_id" validate:"gte=0"`                //父级
    Name             string           `json:"name" validate:"required,max=30,min=4"`     //权限名称
    ApiUrl           string           `json:"api_url" validate:"required,max=200,min=4"` //api地址
    Code             string           `json:"code"  validate:"required,max=50,min=4"`
    PermissionButton PermissionButton `json:"permission_button" validate:"required"` //权限按钮
    PermissionData   PermissionData   `json:"permission_data" validate:"required"`   //权限数据
    Info             string           `json:"info" validate:"max=100"`
    Status           int8             `json:"status" validate:"number,min=0,max=1"` //状态 0=开启 1=关闭
}

type PermissionAddReply struct {
    Id               int              `json:"id"`
    ParentId         int              `json:"parent_id"` //父级
    Name             string           `json:"name"`      //权限名称
    ApiUrl           string           `json:"api_url"`   //api地址
    Code             string           `json:"code"`
    PermissionButton PermissionButton `json:"permission_button"` //权限按钮
    PermissionData   PermissionData   `json:"permission_data"`   //权限数据
    Info             string           `json:"info"`
    Status           StatusValueName  `json:"status"`    //状态 0=开启 1=关闭
    IsDelete         StatusValueName  `json:"is_delete"` //是否删 0=否 1=是
    CreateTime       int64            `json:"create_time"`
}

type PermissionButton struct {
    Post   bool `json:"post" validate:"boolean,isdefault=false"`
    Delete bool `json:"delete" validate:"boolean,isdefault=false"`
    Put    bool `json:"put" validate:"boolean,isdefault=false"`
    Info   bool `json:"info" validate:"boolean,isdefault=false"` //查看详情按钮
}
type PermissionData struct {
    Phone bool `json:"phone" validate:"boolean,isdefault=false"`
}

type PermissionAdminInfoReply struct {
    AdminInfoReply
    Role       []string
    Permission []string
}
