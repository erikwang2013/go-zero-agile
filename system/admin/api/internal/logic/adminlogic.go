package logic

import (
	"context"
	"errors"
	"time"

	dataFormat "erik-agile/common/data-format"
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

func AdminCheckParam(req *types.AdminInfoReq) error {
    validate := validator.New()
    validateRegister(validate)
    var err error
    if req.Id > 0 {
        err = validate.Var(req.Id, "number,max=18,min=1")
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
    if len(req.Phone) > 0 {
        err = validate.Var(req.Phone, "e164")
    }
    if len(req.Email) > 0 {
        err = validate.Var(req.Email, "email")
    }
    if req.Status >= 0 {
        err = validate.Var(req.Email, "number,min=0,max=1,isdefault=-1")
    }
    if req.Gender >= 0 {
        err = validate.Var(req.Email, "number,min=0,max=2,isdefault=-1")
    }
    if req.Page >= 1 {
        err = validate.Var(req.Page, "number,max=11,min=1")
    }
    if req.Limit >= 1 {
        err = validate.Var(req.Limit, "number,max=11,min=1")
    }
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    return nil
}

func (l *AdminLogic) Admin(req *types.AdminInfoReq) (resp []*types.AdminInfoReply, err error) {
    err = AdminCheckParam(req)
    if err != nil {
        return nil, err
    }
    getData := &AdminModel.Admin{
        Id:       req.Id,
        ParentId: req.ParentId,
        Name:     req.Name,
        NickName: req.NickName,
        Phone:    req.Phone,
        Email:    req.Email,
        Gender:   req.Gender,
        Status:   req.Status,
    }
    all, err := l.svcCtx.AdminModel.All(l.ctx, getData, req.Page, req.Limit)
    if err != nil {
        return nil, errors.New("查询用户列表失败")
    }
    var getAll []*types.AdminInfoReply
    for _, v := range all {
        r := &types.AdminInfoReply{
            Id:            int(v.Id),
            ParentId:      v.ParentId,
            HeadImg:       v.HeadImg,
            Name:          v.Name,
            NickName:      v.NickName,
            Password:      "",
            Gender:        types.StatusValueName{Key: v.Gender, Val: AdminModel.AdminGenderName[v.Gender]},
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

    return getAll, nil
}

func (l *AdminLogic) Create(req *types.AdminAddReq) (resp *types.AdminInfoReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    adminInfo, err := l.svcCtx.AdminModel.FindOneName(l.ctx, req.Name)
    if err == nil && adminInfo != nil {
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
        PromotionCode: dataFormat.RandStr(11),
        CreateTime:    getTime,
        UpdateTime:    getTime,
    }
    setData.ParentId = 0
    if req.ParentId >= 1 {
        setData.ParentId = req.ParentId
    }
    password := dataFormat.RandStr(8)
    byct, err := dataFormat.HashAndSalt(password)
    if err != nil {
        return nil, errors.New("密码生成失败")
    }
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
            Key: setData.Gender,
            Val: AdminModel.AdminGenderName[setData.Gender],
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

func (l *AdminLogic) Delete(req *types.AdminInfoReq) (resp *types.AdminInfoReply, err error) {
    return
}

func (l *AdminLogic) Put(req *types.AdminInfoReq) (resp *types.AdminInfoReply, err error) {
    return
}
