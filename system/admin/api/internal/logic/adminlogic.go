package logic

import (
	"context"
	"errors"
	"strings"

	dataFormat "erik-agile/common/data-format"
	"erik-agile/common/date"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"
	"erik-agile/system/admin/model"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
    return &AdminLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *AdminLogic) Create(req *types.AdminAddReq) (code int, resp *types.AdminInfoReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    checkPhone := dataFormat.CheckMobile(req.Phone)
    if len(req.Phone) > 0 && false == checkPhone {
        return 400000, nil, errors.New("手机号格式错误")
    }
    var adminInfo *model.Admin
    resultAdmin := l.svcCtx.Gorm.Model(&model.Admin{}).
        Where("name =? and is_delete=?", req.Name, 0).First(&adminInfo)
    if resultAdmin.RowsAffected > 0 {
        return 400000, nil, errors.New("用户名已存在")
    }
    resultFindPhone := l.svcCtx.Gorm.Model(&model.Admin{}).
        Where("phone=? and is_delete=?", req.Phone, 0).First(&adminInfo)
    if resultFindPhone.RowsAffected > 0 {
        return 400000, nil, errors.New("手机号已存在")
    }
    resultFindEmail := l.svcCtx.Gorm.Model(&model.Admin{}).
        Where("email=? and is_delete=?", req.Email, 0).First(&adminInfo)
    if resultFindEmail.RowsAffected > 0 {
        return 400000, nil, errors.New("邮箱已存在")
    }

    getTime := date.GetDefaultTimeFormat()
    setData := &model.Admin{
        HeadImg:       req.HeadImg,
        Name:          req.Name,
        NickName:      req.NickName,
        Phone:         req.Phone,
        Email:         req.Email,
        Gender:        req.Gender,
        Status:        req.Status,
        Info:          req.Info,
        PromotionCode: dataFormat.RandStr(11),
        CreateTime:    getTime,
        UpdateTime:    getTime,
    }
    setData.ParentId = 0
    if req.ParentId >= 1 {
        setData.ParentId = req.ParentId
    }
    password := dataFormat.RandStr(8)
    if len(req.Password) > 0 {
        password = req.Password
    }
    byct, err := dataFormat.HashAndSalt(password)
    if err != nil {
        return 500000, nil, errors.New("密码生成失败")
    }
    setData.Password = byct
    tx := l.svcCtx.Gorm.Begin()
    resultAdd := tx.Create(&setData)
    if resultAdd.Error != nil {
        tx.Rollback()
        return 500000, nil, errors.New("新增用户失败")
    }
    var roleGroupInfo *model.AdminRoleGroup
    resultFindRole := l.svcCtx.Gorm.Model(&model.AdminRoleGroup{}).
        Where("role_id=? and admin_id=? and is_delete=?", req.RoleId, setData.Id, 0).
        First(&roleGroupInfo)
    if resultFindRole.RowsAffected > 0 {
        tx.Rollback()
        return 400000, nil, errors.New("角色已存分配")
    }
    roleGroup := &model.AdminRoleGroup{
        RoleId:   req.RoleId,
        AdminId:  setData.Id,
        IsDelete: 0,
        Status:   0,
    }
    resultRoleGroup := tx.Create(&roleGroup)
    if resultRoleGroup.Error != nil {
        tx.Rollback()
        return 500000, nil, errors.New("分配用户角色失败")
    }
    tx.Commit()
    getRole, err := getRolePermission(l.svcCtx, setData.Id)
    if err != nil {
        getRole = []*types.RoleAddPermissionReply{}
    }
    return 200000, &types.AdminInfoReply{
        Id:       setData.Id,
        ParentId: setData.ParentId,
        Role:     getRole,
        HeadImg:  setData.HeadImg,
        Name:     setData.Name,
        NickName: setData.NickName,
        Gender: types.StatusValueName{
            Key: setData.Gender,
            Val: model.AdminGenderName[setData.Gender],
        },
        Phone: setData.Phone,
        Email: setData.Email,
        Status: types.StatusValueName{
            Key: setData.Status,
            Val: dataFormat.StatusName[setData.Status],
        },
        IsDelete: types.StatusValueName{
            Key: setData.IsDelete,
            Val: dataFormat.IsDeleteName[setData.IsDelete],
        },
        PromotionCode: setData.PromotionCode,
        Info:          setData.Info,
        CreateTime:    setData.CreateTime.Unix(),
        UpdateTime:    setData.UpdateTime.Unix(),
    }, nil
}

func (l *AdminLogic) Delete(req *types.DeleteIdsReq) (code int, resp *string, err error) {
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
    result := l.svcCtx.Gorm.Model(&model.Admin{}).Where("id IN ?", ids).Updates(model.Admin{IsDelete: 1})
    if result.Error != nil {
        return 500000, nil, errors.New("删除用户失败")
    }
    return 200000, &req.Id, nil
}

func (l *AdminLogic) Put(req *types.AdminPutReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    var up model.Admin
    i := 0
    if req.ParentId > 0 {
        up.ParentId = req.ParentId
        i += 1
    }
    if req.RoleId > 0 {
        i += 1
    }
    if len(req.NickName) > 0 {
        up.NickName = req.NickName
        i += 1
    }
    if len(req.Name) > 0 {
        up.Name = req.Name
        i += 1
    }
    if len(req.Password) > 0 {
        byct, err := dataFormat.HashAndSalt(req.Password)
        if err != nil {
            return 500000, nil, errors.New("密码加密失败")
        }
        up.Password = byct
        i += 1
    }
    if len(req.Phone) > 0 {
        checkPhone := dataFormat.CheckMobile(req.Phone)
        if false == checkPhone {
            return 400000, nil, errors.New("手机号格式错误")
        }
        up.Phone = req.Phone
        i += 1
    }
    if len(req.Email) > 0 {
        up.Email = req.Email
        i += 1
    }
    if req.Status > 0 {
        up.Status = req.Status
        i += 1
    }
    if req.Gender > 0 {
        up.Gender = req.Gender
        i += 1
    }
    if len(req.Info) > 0 {
        up.Info = req.Info
        i += 1
    }
    if i <= 0 {
        return 400000, nil, errors.New("至少更新一个参数")
    }
    var adminInfo *model.Admin
    resultAdmin := l.svcCtx.Gorm.Model(&model.Admin{}).
        Where("id <> ? and name=? and is_delete=?", req.Id, req.Name, 0).
        First(&adminInfo)
    if resultAdmin.RowsAffected > 0 {
        return 400000, nil, errors.New("用户名已存在")
    }
    resultFindPhone := l.svcCtx.Gorm.Model(&model.Admin{}).
        Where("id <> ? and phone=? and is_delete=?", req.Id, req.Phone, 0).
        First(&adminInfo)
    if resultFindPhone.RowsAffected > 0 {
        return 400000, nil, errors.New("手机号已存在")
    }
    resultFindEmail := l.svcCtx.Gorm.Model(&model.Admin{}).
        Where("id <> ? and email=? and is_delete=?", req.Id, req.Email, 0).
        First(&adminInfo)
    if resultFindEmail.RowsAffected > 0 {
        return 400000, nil, errors.New("邮箱已存在")
    }
    tx := l.svcCtx.Gorm.Begin()
    result := tx.Model(&model.Admin{}).Where("id = ?", req.Id).Updates(up)
    if result.Error != nil {
        tx.Rollback()
        return 500000, nil, errors.New("更新用户失败")
    }
    var roleGroupInfo *model.AdminRoleGroup
    resultFindRole := l.svcCtx.Gorm.Model(&model.AdminRoleGroup{}).
        Where("role_id=? and admin_id=? and is_delete=?", req.RoleId, req.Id, 0).
        First(&roleGroupInfo)
    if resultFindRole.RowsAffected > 0 {
        tx.Rollback()
        return 400000, nil, errors.New("角色已存分配")
    }
    roleGroup := &model.AdminRoleGroup{
        RoleId:   req.RoleId,
        AdminId:  req.Id,
        IsDelete: 0,
        Status:   0,
    }
    resultRoleGroup := tx.Create(&roleGroup)
    if resultRoleGroup.Error != nil {
        tx.Rollback()
        return 500000, nil, errors.New("分配用户角色失败")
    }
    tx.Commit()
    upId := dataFormat.IntToString(req.Id)
    return 200000, &upId, nil
}

func AdminCheckParam(req *types.AdminSearchReq) error {
    validate := validator.New()
    validateRegister(validate)
    var err error
    if req.Id > 0 {
        err = validate.Var(req.Id, "gte=0")
    }
    if req.ParentId >= 0 {
        err = validate.Var(req.ParentId, "number,max=18,min=0,isdefault=-1")
    }
    if len(req.NickName) > 0 {
        err = validate.Var(req.NickName, "alphanum,max=30,min=4")
    }
    if len(req.Name) > 0 {
        err = validate.Var(req.Name, "alphanum,max=30,min=4")
    }
    checkPhone := dataFormat.CheckMobile(req.Phone)
    if len(req.Phone) > 0 && false == checkPhone {
        return errors.New("手机号格式错误")
    }
    if len(req.Email) > 0 {
        err = validate.Var(req.Email, "email")
    }
    if req.Status >= 0 {
        err = validate.Var(req.Status, "number,oneof=-1 0 1")
    }
    if req.Gender >= 0 {
        err = validate.Var(req.Gender, "number,oneof=-1 0 1 2")
    }
    if req.Page >= 1 {
        err = validate.Var(req.Page, "number,lte=10000,gte=1")
    }
    if req.Limit >= 1 {
        err = validate.Var(req.Limit, "number,lte=1000,gte=1")
    }
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    return nil
}

func (l *AdminLogic) Index(req *types.AdminSearchReq) (code int, resp []*types.AdminInfoReply, err error) {
    err = AdminCheckParam(req)
    if err != nil {
        return 400000, nil, err
    }
    var all []*model.Admin
    var getData model.Admin
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
    if len(req.NickName) > 0 {
        getData.NickName = req.NickName
    }
    if len(req.Phone) > 0 {
        getData.Phone = req.Phone
    }
    if len(req.Email) > 0 {
        getData.Email = req.Email
    }
    if req.Status >= 0 {
        getData.Status = req.Status
    }
    if req.Gender >= 0 {
        getData.Gender = req.Gender
    }
    if req.Limit <= 0 {
        req.Limit = 10
    }
    if req.Page <= 0 {
        req.Page = 1
    }
    var total int64
    db := l.svcCtx.Gorm.Model(&model.Admin{}).Where(&getData)
    db.Count(&total)
    pageSetNum, offset := dataFormat.Page(req.Limit, req.Page, total)
    result := db.Limit(pageSetNum).Offset(offset).Find(&all)
    if result.Error != nil {
        return 500000, nil, errors.New("查询用户列表失败")
    }
    getAll := []*types.AdminInfoReply{}
    if len(all) <= 0 {
        return 200000, getAll, nil
    }

    for _, v := range all {
        getRole, err := getRolePermission(l.svcCtx, v.Id)
        if err != nil {
            getRole = []*types.RoleAddPermissionReply{}
        }
        r := &types.AdminInfoReply{
            Id:            int(v.Id),
            ParentId:      v.ParentId,
            Role:          getRole,
            HeadImg:       v.HeadImg,
            Name:          v.Name,
            NickName:      v.NickName,
            Gender:        types.StatusValueName{Key: v.Gender, Val: model.AdminGenderName[v.Gender]},
            Phone:         v.Phone,
            Email:         v.Email,
            Status:        types.StatusValueName{Key: v.Status, Val: dataFormat.StatusName[v.Status]},
            IsDelete:      types.StatusValueName{Key: v.IsDelete, Val: dataFormat.IsDeleteName[v.IsDelete]},
            PromotionCode: v.PromotionCode,
            Info:          v.Info,
            CreateTime:    v.CreateTime.Unix(),
            UpdateTime:    v.UpdateTime.Unix(),
        }
        getAll = append(getAll, r)
    }

    return 200000, getAll, nil
}

//获取个人信息
func (l *AdminLogic) AdminInfo(req *types.AdminInfoAllReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    if req.Id > 0 {
        err = validate.Var(req.Id, "required,gte=0")
        if err != nil {
            varError := err.(validator.ValidationErrors)
            transStr := varError.Translate(trans)
            return 400001, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
        }
    }

    return 200000, nil, nil
}
