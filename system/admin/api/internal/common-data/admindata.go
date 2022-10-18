package commonData

import (
	"context"
	dataFormat "erik-agile/common/data-format"
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/types"
	"erik-agile/system/admin/model"
	"errors"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

//校验权限
func CheckPermission(Gorm *gorm.DB, ctx context.Context, c config.Config, url, method string) bool {
    checkStr := dataFormat.GetMd5(url + method)
    result, err := GetRolePermission(Gorm, ctx)
    if err != nil {
        logx.Error("校验权限异常")
        logx.Error(err)
        return false
    }

    if len(result) <= 0 {
        logx.Error("账户未配置权限")
        return false
    }
    roleAll := c.Permission.Role
    rolePermission := 0
    permissionUrl := map[string]bool{}
    for _, v := range result {
        if strings.Compare(v.Code, roleAll) == 0 {
            rolePermission = 1
            break
        }
        if len(v.Permission) <= 0 {
            continue
        }
        for _, p := range v.Permission {
            permissionUrl[p.Code] = true
        }
    }
    if rolePermission == 1 {
        return true
    }
    if len(permissionUrl) <= 0 {
        logx.Error("账户权限不存在")
        return false
    }
    if permissionUrl[checkStr] == true {
        return true
    }
    logx.Error("账户校验完成，校验异常")
    return false
}

//获取用户id
func GetAdminId(ctx context.Context) int {
    adminId := ctx.Value("admin_id")
    getAdminId := fmt.Sprintf("%v", adminId)
    return dataFormat.StringToInt(getAdminId)
}

func GetRolePermissionArr(Gorm *gorm.DB, ctx context.Context) []*types.RoleAddPermissionReply {
    result, err := GetRolePermission(Gorm, ctx)
    if err != nil {
        return []*types.RoleAddPermissionReply{}
    }
    return result
}

//获取用户的角色及权限
func GetRolePermission(Gorm *gorm.DB, ctx context.Context, getConfig ...string) (resp []*types.RoleAddPermissionReply, err error) {
    var all []*model.AdminRoleGroup
    adminId := GetAdminId(ctx)
    result := Gorm.Model(&model.AdminRoleGroup{}).
        Where("admin_id = ? AND is_delete= ?", adminId, 0).Find(&all)

    if result.Error != nil {
        return nil, errors.New("用户组不存在")
    }
    var roleIds []int
    for _, v := range all {
        roleIds = append(roleIds, v.RoleId)
    }
    if len(roleIds) <= 0 {
        return nil, errors.New("用户角色不存在")
    }

    //查询角色
    var allRole []*model.Role
    resultRole := Gorm.Model(&model.Role{}).
        Where("id IN ? AND is_delete= ?", roleIds, 0).
        Find(&allRole)
    if resultRole.Error != nil {
        return nil, errors.New("角色不存在")
    }
    getRoleAll, err := GetRole(Gorm, allRole)
    if err != nil {
        return nil, errors.New("用户角色权限不存在")
    }
    return getRoleAll, nil
}

func GetRole(Gorm *gorm.DB, allRole []*model.Role) (resp []*types.RoleAddPermissionReply, err error) {
    getRoleAll := []*types.RoleAddPermissionReply{}
    if len(allRole) <= 0 {
        return getRoleAll, nil
    }
    roleIds := []int{}
    for _, v := range allRole {
        roleIds = append(roleIds, v.Id)
    }
    var allRolePermission []*model.RolePermission
    rolePermission := Gorm.Model(&model.RolePermission{}).
        Where("role_id in ? AND  is_delete= ?", roleIds, 0).Find(&allRolePermission)
    if rolePermission.Error != nil {
        return nil, errors.New("角色权限不存在")
    }
    if len(allRolePermission) <= 0 {
        return nil, errors.New("角色权限未设置")
    }
    var perIds []int
    for _, vp := range allRolePermission {
        perIds = append(perIds, vp.PermissionId)
    }
    var allPermission []*model.Permission
    resultRole := Gorm.Model(&model.Permission{}).
        Where("id IN ? and is_delete = ?", perIds, 0).
        Find(&allPermission)
    if resultRole.Error != nil {
        return nil, errors.New("权限不存在")
    }

    if len(allPermission) <= 0 {
        return nil, errors.New("权限未设置")
    }
    for _, v := range allRole {
        getAllPer := []*types.PermissionGetReply{}
        for _, pr := range allRolePermission {
            if v.Id != pr.RoleId {
                continue
            }
            for _, vpd := range allPermission {
                if pr.PermissionId == vpd.Id {
                    rpd := &types.PermissionGetReply{
                        Id:       int(vpd.Id),
                        ParentId: vpd.ParentId,
                        ApiUrl:   vpd.ApiUrl,
                        Method:   vpd.Method,
                        Name:     vpd.Name,
                        Status:   types.StatusValueName{Key: vpd.Status, Val: dataFormat.StatusName[vpd.Status]},
                        Code:     vpd.Code,
                    }
                    getAllPer = append(getAllPer, rpd)
                }
            }
        }
        r := &types.RoleAddPermissionReply{
            Id:         v.Id,
            ParentId:   v.ParentId,
            Name:       v.Name,
            Status:     types.StatusValueName{Key: v.Status, Val: dataFormat.StatusName[v.Status]},
            Code:       v.Code,
            Permission: getAllPer,
        }
        getRoleAll = append(getRoleAll, r)
    }
    return getRoleAll, nil
}
