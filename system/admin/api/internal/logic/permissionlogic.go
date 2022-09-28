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

type PermissionLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionLogic {
    return &PermissionLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *PermissionLogic) Create(req *types.PermissionAddReq) (code int, resp *types.PermissionAddReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    var findData *model.Permission
    resultFindCode := l.svcCtx.Gorm.Model(&model.Permission{}).Where(&model.Permission{Code: req.Code}).First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errors.New("权限编码已存在")
    }

    resultFindUrl := l.svcCtx.Gorm.Model(&model.Permission{}).
        Where(&model.Permission{ApiUrl: req.ApiUrl, Method: req.Method}).First(&findData)
    if resultFindUrl.RowsAffected > 0 {
        return 400000, nil, errors.New("url和请求类型已存在")
    }

    setData := &model.Permission{
        ParentId:   req.ParentId,
        Name:       req.Name,
        ApiUrl:     req.ApiUrl,
        Code:       req.Code,
        Status:     req.Status,
        Method:     req.Method,
        IsDelete:   0,
        CreateTime: date.GetDefaultTimeFormat(),
    }
    if len(req.Info) > 0 {
        setData.Info = req.Info
    }
    result := l.svcCtx.Gorm.Create(&setData)
    if result.Error != nil {
        return 500000, nil, errors.New("新增权限失败")
    }
    return 200000, &types.PermissionAddReply{
        Id:       setData.Id,
        ParentId: setData.ParentId,
        Name:     setData.Name,
        ApiUrl:   setData.ApiUrl,
        Method:   setData.Method,
        Code:     setData.Code,
        Status: types.StatusValueName{
            Key: setData.Status,
            Val: dataFormat.StatusName[req.Status],
        },
        IsDelete: types.StatusValueName{
            Key: setData.IsDelete,
            Val: dataFormat.IsDeleteName[setData.IsDelete],
        },
        Info:       setData.Info,
        CreateTime: setData.CreateTime.Unix(),
    }, nil
}

func (l *PermissionLogic) Delete(req *types.DeleteIdsReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    var ids []string
    if len(req.Id) <= 0 {
        return 400000, nil, errors.New("删除id必须")
    }
    ids = strings.Split(req.Id, ",")
    for _, v := range ids {
        err = validate.Var(v, "alphanum,max=100,min=1")
        if err != nil {
            varError := err.(validator.ValidationErrors)
            transStr := varError.Translate(trans)
            return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
        }
    }
    result := l.svcCtx.Gorm.Model(&model.Permission{}).Where("id IN ?", ids).Updates(model.Permission{IsDelete: 1})
    if result.Error != nil {
        return 500000, nil, errors.New("删除权限失败")
    }
    return 200000, &req.Id, nil
}

func (l *PermissionLogic) Put(req *types.PermissionPutReq) (code int, resp *string, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    var up model.Permission
    i := 0
    if req.ParentId > 0 {
        up.ParentId = req.ParentId
        i += 1
    }
    if len(req.Name) > 0 {
        up.Name = req.Name
        i += 1
    }
    if len(req.ApiUrl) > 0 {
        up.ApiUrl = req.ApiUrl
        i += 1
    }
    if len(req.Method) > 0 {
        up.Method = req.Method
        i += 1
    }
    if len(req.Code) > 0 {
        up.Code = req.Code
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
    var findData *model.Permission
    resultFindCode := l.svcCtx.Gorm.Model(&model.Permission{}).
        Where("id <> ?", req.Id).
        Where(&model.Permission{Code: req.Code}).First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errors.New("权限编码已存在")
    }

    resultFindUrl := l.svcCtx.Gorm.Model(&model.Permission{}).
        Where("id <> ?", req.Id).
        Where(&model.Permission{ApiUrl: req.ApiUrl, Method: req.Method}).First(&findData)
    if resultFindUrl.RowsAffected > 0 {
        return 400000, nil, errors.New("url和请求类型已存在")
    }
    result := l.svcCtx.Gorm.Model(&model.Permission{}).Where("id = ?", req.Id).Updates(up)
    if result.Error != nil {
        return 500000, nil, errors.New("更新用户失败")
    }
    upId := dataFormat.IntToString(req.Id)
    return 200000, &upId, nil
}

func (l *PermissionLogic) Permission(req *types.AdminInfoReq) (code int, resp []*types.AdminInfoReply, err error) {
    return 200000, nil, nil
}
