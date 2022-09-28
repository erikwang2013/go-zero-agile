package logic

import (
	"context"
	"encoding/json"
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
    err = validate.Struct(req.PermissionButton)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    err = validate.Struct(req.PermissionData)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    setData := model.Permission{
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
    buttonJson, _ := json.Marshal(req.PermissionButton)
    setData.PermissionButton = string(buttonJson)
    pDataJson, _ := json.Marshal(req.PermissionData)
    setData.PermissionData = string(pDataJson)
    result := l.svcCtx.Gorm.Create(&setData)
    if result.Error != nil {
        return 500000, nil, errors.New("新增权限失败")
    }
    var getButton types.PermissionButton
    var getdata types.PermissionData
    err = json.Unmarshal([]byte(setData.PermissionButton), &getButton)
    if err != nil {
        return 500000, nil, errors.New("解析按钮权限错误")
    }
    err = json.Unmarshal([]byte(setData.PermissionData), &getdata)
    if err != nil {
        return 500000, nil, errors.New("解析数据权限错误")
    }
    return 200000, &types.PermissionAddReply{
        Id:               setData.Id,
        ParentId:         setData.ParentId,
        Name:             setData.Name,
        ApiUrl:           setData.ApiUrl,
        Code:             setData.Code,
        PermissionButton: getButton,
        PermissionData:   getdata,
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
