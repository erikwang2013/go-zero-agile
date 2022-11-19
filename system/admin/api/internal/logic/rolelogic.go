package logic

import (
	"context"
	dataFormat "erik-agile/common/data-format"
	"erik-agile/common/date"
	"erik-agile/common/errorx"
	commonData "erik-agile/system/admin/api/internal/common-data"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/svc/gorm"
	"erik-agile/system/admin/api/internal/types"
	"erik-agile/system/admin/model"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
)

type RoleLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
    db     *gorm.Gormdb
}

func NewRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleLogic {
    return &RoleLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
        db:     gorm.NewGormdb(svcCtx.Config),
    }
}

func (l *RoleLogic) Create(req *types.RoleAddReq) (code int, resp *types.RoleAddReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errorx.NewDefaultError(dataFormat.RemoveTopStruct(transStr))
    }
    CheckCode := dataFormat.GetMd5(stringx.Rand())
    var findData *model.Role
    resultFindCode := l.db.Gorm.Where(&model.Role{Code: CheckCode}).First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errorx.NewDefaultError("角色编码已存在")
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
    tx := l.db.Gorm.Begin()
    result := tx.Create(&setData)
    if result.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("新增角色失败")
    }
    permissionIds := strings.Split(req.Permission, ",")
    rolePermission := []model.RolePermission{}
    if len(permissionIds) <= 0 {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("新增角色权限失败")
    }
    for _, pr := range permissionIds {
        rolePermission = append(rolePermission, model.RolePermission{
            RoleId:              setData.Id,
            PermissionId:        dataFormat.StringToInt(pr),
            PermissionChilderId: "[]",
            PermissionButton:    "[]",
            Status:              0,
            IsDelete:            0,
        })
    }
    resultPr := tx.Create(&rolePermission)
    if resultPr.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("新增角色权限失败")
    }
    tx.Commit()
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
        return 400000, nil, errorx.NewDefaultError("删除id必须")
    }
    ids = strings.Split(req.Id, ",")
    for _, v := range ids {
        err = validate.Var(v, "alphanum,gte=1")
        if err != nil {
            varError := err.(validator.ValidationErrors)
            transStr := varError.Translate(trans)
            return 400000, nil, errorx.NewDefaultError(dataFormat.RemoveTopStruct(transStr))
        }
    }
    tx := l.db.Gorm.Begin()
    result := tx.Model(&model.Role{}).Where("id IN ?", ids).Updates(model.Role{IsDelete: 1})
    if result.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("删除角色失败")
    }
    resultPr := tx.Model(&model.RolePermission{}).Where("role_id IN ?", ids).Updates(model.RolePermission{IsDelete: 1})
    if resultPr.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("删除角色权限失败")
    }
    tx.Commit()
    return 200000, &req.Id, nil
}

func (l *RoleLogic) Put(req *types.RolePutReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errorx.NewDefaultError(dataFormat.RemoveTopStruct(transStr))
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
        return 400000, nil, errorx.NewDefaultError("至少更新一个参数")
    }
    CheckCode := dataFormat.GetMd5(stringx.Rand())
    up.Code = CheckCode
    var findData *model.Role
    resultFindCode := l.db.Gorm.Model(&model.Role{}).
        Where("id <> ? and code=?", req.Id, CheckCode).
        First(&findData)
        
    if resultFindCode.RowsAffected >0 {
        return 400000, nil, errorx.NewDefaultError("角色编码已存在")
    }
    tx := l.db.Gorm.Begin()
    result := tx.Model(&model.Role{}).Where("id = ?", req.Id).Updates(up)
    if result.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("更新角色失败")
    }
    resultPr := tx.Model(&model.RolePermission{}).Where("role_id = ?", req.Id).Updates(model.RolePermission{IsDelete: 1})
    if resultPr.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("删除角色权限失败")
    }
    permissionIds := strings.Split(req.Permission, ",")
    rolePermission := []model.RolePermission{}
    if len(permissionIds) <= 0 {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("新增角色权限失败")
    }
    for _, pr := range permissionIds {
        rolePermission = append(rolePermission, model.RolePermission{
            RoleId:              req.Id,
            PermissionId:        dataFormat.StringToInt(pr),
            PermissionChilderId: "[]",
            PermissionButton:    "[]",
            Status:              0,
            IsDelete:            0,
        })
    }
    resultPrC := tx.Create(&rolePermission)
    if resultPrC.Error != nil {
        tx.Rollback()
        return 500000, nil, errorx.NewDefaultError("新增角色权限失败")
    }
    tx.Commit()
    upId := dataFormat.IntToString(req.Id)
    return 200000, &upId, nil
}

func (l *RoleLogic) Index(req *types.RoleSearchReq) (code int, resp []*types.RoleAddPermissionReply, err error) {
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
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errorx.NewDefaultError(dataFormat.RemoveTopStruct(transStr))
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
    db := l.db.Gorm.Model(&model.Role{}).Where(&getData)
    if req.ParentId >= 0 {
        db = db.Where("parent_id =?", req.ParentId)
    }
    if req.Status >= 0 {
        db = db.Where("status =?", req.Status)
    }
    result := db.Find(&all)
    if result.Error != nil {
        return 500000, nil, errorx.NewDefaultError("查询角色列表失败")
    }
    getAll := []*types.RoleAddPermissionReply{}
    if len(all) <= 0 {
        return 404000, getAll, errorx.NewDefaultError("角色不存在或异常")
    }
    getAll, err = commonData.GetRole(l.db.Gorm, all)
    if err != nil {
        return 500000, nil, errorx.NewDefaultError("角色不存在或异常")
    }
    // for _, v := range permission {
    //     r := &types.RoleAddReply{
    //         Id:         v.Id,
    //         ParentId:   v.ParentId,
    //         Name:       v.Name,
    //         Status:     types.StatusValueName{Key: v.Status, Val: dataFormat.StatusName[v.Status]},
    //         IsDelete:   types.StatusValueName{Key: v.IsDelete, Val: dataFormat.IsDeleteName[v.IsDelete]},
    //         Code:       v.Code,
    //         Info:       v.Info,
    //         CreateTime: v.CreateTime.Unix(),
    //     }
    //     getAll = append(getAll, r)
    // }
    return 200000, getAll, nil
}
