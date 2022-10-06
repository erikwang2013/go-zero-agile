package logic

import (
	"context"
	dataFormat "erik-agile/common/data-format"
	"erik-agile/common/date"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"
	"erik-agile/system/admin/model"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type RoleLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
    return &RoleLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

//获取用户的角色及权限
func getRolePermission(svcCtx *svc.ServiceContext, adminId int) (resp []*types.RoleAddPermissionReply, err error) {
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
            Id:         int(v.Id),
            ParentId:   v.ParentId,
            Name:       v.Name,
            Status:     types.StatusValueName{Key: v.Status, Val: dataFormat.StatusName[v.Status]},
            Code:       v.Code,
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
                Id:         int(vpd.Id),
                ParentId:   vpd.ParentId,
                ApiUrl:     vpd.ApiUrl,
                Method:     vpd.Method,
                Name:       vpd.Name,
                Status:     types.StatusValueName{Key: vpd.Status, Val: dataFormat.StatusName[vpd.Status]},
                Code:       vpd.Code,
            }
            getAllPer = append(getAllPer, rpd)
        }
        r.Permission = getAllPer
        getRoleAll = append(getRoleAll, r)
    }
    return getRoleAll, nil
}

func (l *RoleLogic) Create(req *types.RoleAddReq) (code int, resp *types.RoleAddReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    CheckCode := dataFormat.GetMd5(stringx.Rand())
    var findData *model.Role
    resultFindCode := l.svcCtx.Gorm.Model(&model.Role{}).Where(&model.Role{Code: CheckCode}).First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errors.New("角色编码已存在")
    }

    setData := &model.Role{
        ParentId:   req.ParentId,
        Name:       req.Name,
        Code:       CheckCode,
        Status:     req.Status,
        IsDelete:   0,
        CreateTime: date.GetDefaultTimeFormat(),
    }
    if len(req.Info) > 0 {
        setData.Info = req.Info
    }
    result := l.svcCtx.Gorm.Create(&setData)
    if result.Error != nil {
        return 500000, nil, errors.New("新增角色失败")
    }
    return 200000, &types.RoleAddReply{
        Id:       setData.Id,
        ParentId: setData.ParentId,
        Name:     setData.Name,
        Code:     setData.Code,
        Status: types.StatusValueName{
            Key: setData.Status,
            Val: dataFormat.StatusName[setData.Status],
        },
        IsDelete: types.StatusValueName{
            Key: setData.IsDelete,
            Val: dataFormat.IsDeleteName[setData.IsDelete],
        },
        Info:       setData.Info,
        CreateTime: setData.CreateTime.Unix(),
    }, nil
}

func (l *RoleLogic) Delete(req *types.DeleteIdsReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    var ids []string
    if len(req.Id) <= 0 {
        return 400000, nil, errors.New("删除id必须")
    }
    ids = strings.Split(req.Id, ",")
    for _, v := range ids {
        err = validate.Var(v, "alphanum,gte=1")
        if err != nil {
            varError := err.(validator.ValidationErrors)
            transStr := varError.Translate(trans)
            return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
        }
    }
    result := l.svcCtx.Gorm.Model(&model.Role{}).Where("id IN ?", ids).Updates(model.Role{IsDelete: 1})
    if result.Error != nil {
        return 500000, nil, errors.New("删除权限失败")
    }
    return 200000, &req.Id, nil
}

func (l *RoleLogic) Put(req *types.RolePutReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    var up model.Role
    i := 0
    if req.ParentId > 0 {
        up.ParentId = req.ParentId
        i += 1
    }
    if len(req.Name) > 0 {
        up.Name = req.Name
        i += 1
    }
    if len(req.Info) > 0 {
        up.Info = req.Info
        i += 1
    }
    if req.Status > 0 {
        up.Status = req.Status
        i += 1
    }
    if i <= 0 {
        return 400000, nil, errors.New("至少更新一个参数")
    }
    CheckCode := dataFormat.GetMd5(stringx.Rand())
    up.Code = CheckCode
    var findData *model.Role
    resultFindCode := l.svcCtx.Gorm.Model(&model.Role{}).
        Where("id <> ?", req.Id).
        Where(&model.Role{Code: CheckCode}).First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errors.New("角色编码已存在")
    }
    result := l.svcCtx.Gorm.Model(&model.Role{}).Where("id = ?", req.Id).Updates(up)
    if result.Error != nil {
        return 500000, nil, errors.New("更新角色失败")
    }
    upId := dataFormat.IntToString(req.Id)
    return 200000, &upId, nil
}

func (l *RoleLogic) Index(req *types.RoleSearchReq) (code int, resp []*types.RoleAddReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    if req.Id > 0 {
        err = validate.Var(req.Id, "gt=0")
    }
    if req.ParentId >= 0 {
        err = validate.Var(req.ParentId, "number,gte=0")
    }
    if len(req.Name) > 0 {
        err = validate.Var(req.Name, "alphanum,max=30,min=4")
    }
    if len(req.Code) > 0 {
        err = validate.Var(req.Code, "max=50,min=4")
    }
    if req.Status >= 0 {
        err = validate.Var(req.Status, "oneof=-1 0 1")
    }
    // if req.Page >= 1 {
    //     err = validate.Var(req.Page, "number,lte=10000,gte=1")
    // }
    // if req.Limit >= 1 {
    //     err = validate.Var(req.Limit, "number,lte=1000,gte=1")
    // }
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    var getData model.Role
    getData.IsDelete = int8(0)
    if req.Id > 0 {
        getData.Id = req.Id
    }
    if req.ParentId >= 0 {
        getData.ParentId = req.ParentId
    }
    if len(req.Name) > 0 {
        getData.Name = req.Name
    }
    if len(req.Code) > 0 {
        getData.Code = req.Code
    }
    if req.Status >= 0 {
        getData.Status = req.Status
    }
    var all []*model.Role
    //var total int64
    db := l.svcCtx.Gorm.Model(&model.Role{}).Where(&getData)
    // db.Count(&total)
    // pageSetNum, offset := dataFormat.Page(req.Limit, req.Page, total)
    //result := db.Limit(pageSetNum).Offset(offset).Find(&all)
    result := db.Find(&all)
    if result.Error != nil {
        return 500000, nil, errors.New("查询用户列表失败")
    }
    getAll := []*types.RoleAddReply{}
    for _, v := range all {
        r := &types.RoleAddReply{
            Id:         int(v.Id),
            ParentId:   v.ParentId,
            Name:       v.Name,
            Status:     types.StatusValueName{Key: v.Status, Val: dataFormat.StatusName[v.Status]},
            IsDelete:   types.StatusValueName{Key: v.IsDelete, Val: dataFormat.IsDeleteName[v.IsDelete]},
            Code:       v.Code,
            Info:       v.Info,
            CreateTime: v.CreateTime.Unix(),
        }
        getAll = append(getAll, r)
    }
    return 200000, getAll, nil
}
