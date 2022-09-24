package logic

import (
	"context"
	"errors"
	"time"

	"erik-agile/common/data"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"
	AdminModel "erik-agile/system/admin/model/admin"

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

func (l *AdminLogic) Admin(req *types.AdminInfoReq) (resp *types.AdminInfoReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return nil, errors.New(data.RemoveTopStruct(transStr))
    }
    return
}

func (l *AdminLogic) Create(req *types.AdminAddReq) (resp *types.AdminInfoReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return nil, errors.New(data.RemoveTopStruct(transStr))
    }
    adminInfo, err := l.svcCtx.AdminModel.FindOneName(l.ctx, req.Name)
    if err == nil && adminInfo != nil{
        return nil, errors.New("用户名已存在")
    }
    getTime := time.Unix(time.Now().Unix(), 0)
    setData := &AdminModel.Admin{
        HeadImg:       req.HeadImg,
        Name:          req.Name,
        NickName:      req.NickName,
        Phone:         req.Phone,
        Email:         req.Email,
        Gender:        req.Gender,
        Status:        req.Status,
        Info:          req.Info,
        PromotionCode: data.RandStr(11),
        CreateTime:    getTime,
        UpdateTime:    getTime,
    }
    setData.ParentId = 0
    if req.ParentId >= 1 {
        setData.ParentId = req.ParentId
    }
    password := data.RandStr(8)
    logx.Info("===密码生成=1=")
    logx.Info(password)
    byct, err := data.HashAndSalt(password)
    if err != nil {
        return nil, errors.New("密码生成失败")
    }
    logx.Info("===密码生成=2=")
    logx.Info(byct)
    setData.Password = byct
    insert, err := l.svcCtx.AdminModel.Insert(l.ctx, setData)
    if err != nil {
        return nil, errors.New("新增用户失败")
    }
    getId, _ := insert.LastInsertId()
    return &types.AdminInfoReply{
        Id:       int(getId),
        ParentId: setData.ParentId,
        HeadImg:  setData.HeadImg,
        Name:     setData.Name,
        NickName: setData.NickName,
        Password: password,
        Gender: types.StatusValueName{
            Id:   setData.Gender,
            Name: AdminModel.AdminGenderName[setData.Gender],
        },
        Phone: setData.Phone,
        Email: setData.Email,
        Status: types.StatusValueName{
            Id:   setData.Status,
            Name: data.StatusName[setData.Status],
        },
        IsDelete: types.StatusValueName{
            Id:   setData.IsDelete,
            Name: data.IsDeleteName[setData.IsDelete],
        },
        PromotionCode: setData.PromotionCode,
        Info:          setData.Info,
        CreateTime:    setData.CreateTime.Unix(),
        UpdateTime:    setData.UpdateTime.Unix(),
    }, nil
}

func (l *AdminLogic) Delete(req *types.AdminInfoReq) (resp *types.AdminInfoReply, err error) {
    return
}

func (l *AdminLogic) Put(req *types.AdminInfoReq) (resp *types.AdminInfoReply, err error) {
    return
}
