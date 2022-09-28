package logic

import (
	"context"
	"errors"

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
    setData := &model.Permission{
        ParentId:   req.ParentId,
        Name:       req.Name,
        ApiUrl:     req.ApiUrl,
        Code:       req.Code,
        Status:     req.Status,
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

func (l *PermissionLogic) Delete(req *types.AdminInfoReq) (code int, resp []*types.AdminInfoReply, err error) {
    return 200000, nil, nil
}

func (l *PermissionLogic) Put(req *types.AdminInfoReq) (code int, resp []*types.AdminInfoReply, err error) {
    return 200000, nil, nil
}

func (l *PermissionLogic) Permission(req *types.AdminInfoReq) (code int, resp []*types.AdminInfoReply, err error) {
    return 200000, nil, nil
}
