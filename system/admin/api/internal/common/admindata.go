package common

import (
    "context"
    dataFormat "erik-agile/common/data-format"
    "erik-agile/system/admin/api/internal/svc"
    "erik-agile/system/admin/api/internal/types"
    "erik-agile/system/admin/model"
    "errors"
    "fmt"
)

//校验权限
func CheckPermission(url, method string) bool {
    checkStr := dataFormat.GetMd5(url + method)
    logx.Error(checkStr)
    return true
}

//获取用户id
func GetAdminId(ctx context.Context) int {
    adminId := ctx.Value("admin_id")
    getAdminId := fmt.Sprintf("%v", adminId)
    return dataFormat.StringToInt(getAdminId)
}

//获取用户的角色及权限
func GetRolePermission(svcCtx *svc.ServiceContext, adminId int) (resp []*types.RoleAddPermissionReply, err error) {
    var all []*model.AdminRoleGroup
    result := svcCtx.Gorm.Model(&model.AdminRoleGroup{}).
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
    resultRole := svcCtx.Gorm.Model(&model.Role{}).
        Where("id IN ? AND is_delete= ?", roleIds, 0).
        Find(&allRole)
    if resultRole.Error != nil {
        return nil, errors.New("角色不存在")
    }
    getRoleAll := []*types.RoleAddPermissionReply{}
    if len(allRole) <= 0 {
        return getRoleAll, nil
    }
    for _, v := range allRole {
        r := &types.RoleAddPermissionReply{
            Id:       int(v.Id),
            ParentId: v.ParentId,
            Name:     v.Name,
            Status:   types.StatusValueName{Key: v.Status, Val: dataFormat.StatusName[v.Status]},
            Code:     v.Code,
        }
        var allRolePermission []*model.RolePermission
        rolePermission := svcCtx.Gorm.Model(&model.RolePermission{}).
            Where("role_id = ? AND  is_delete= ?", v.Id, 0).Find(&allRolePermission)
        if rolePermission.Error != nil {
            continue
            //return nil, errors.New("角色权限不存在")
        }
        var perIds []int
        if len(allRolePermission) <= 0 {
            continue
        }
        for _, vp := range allRolePermission {
            perIds = append(perIds, vp.PermissionId)
        }
        var allPermission []*model.Permission
        resultRole := svcCtx.Gorm.Model(&model.Permission{}).
            Where("id IN ? and is_delete = ?", perIds, 0).
            Find(&allPermission)
        if resultRole.Error != nil {
            continue
            //return nil, errors.New("权限不存在")
        }
        getAllPer := []*types.PermissionGetReply{}
        if len(allPermission) <= 0 {
            continue
        }
        for _, vpd := range allPermission {
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
        r.Permission = getAllPer
        getRoleAll = append(getRoleAll, r)
    }
    return getRoleAll, nil
}
