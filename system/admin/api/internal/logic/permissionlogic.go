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
    CheckCode := dataFormat.GetMd5(req.ApiUrl + req.Method)
    var findData *model.Permission
    resultFindCode := l.svcCtx.Gorm.Model(&model.Permission{}).
        Where(&model.Permission{Code: CheckCode, IsDelete: 0}).
        First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errors.New("权限编码已存在")
    }

    resultFindUrl := l.svcCtx.Gorm.
        Where(&model.Permission{ApiUrl: req.ApiUrl, Method: req.Method}).
        First(&findData)
    if resultFindUrl.RowsAffected > 0 {
        return 400000, nil, errors.New("url和请求类型已存在")
    }

    setData := &model.Permission{
        ParentId:   req.ParentId,
        Name:       req.Name,
        ApiUrl:     req.ApiUrl,
        Status:     req.Status,
        Method:     req.Method,
        Code:       CheckCode,
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

func (l *PermissionLogic) Delete(req *types.DeleteIdsReq) (code int, resp *string, err error) {
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
    CheckCode := dataFormat.GetMd5(req.ApiUrl + req.Method)
    up.Code = CheckCode
    var findData *model.Permission
    resultFindCode := l.svcCtx.Gorm.Model(&model.Permission{}).
        Where("id <> ? and code=? and is_delete=?", req.Id, CheckCode, 0).
        First(&findData)
    if resultFindCode.RowsAffected > 0 {
        return 400000, nil, errors.New("权限编码已存在")
    }

    resultFindUrl := l.svcCtx.Gorm.Model(&model.Permission{}).
        Where("id <> ? and api_url=? and method=? and is_delete=?", req.Id, req.ApiUrl, req.Method, 0).
        First(&findData)
    if resultFindUrl.RowsAffected > 0 {
        return 400000, nil, errors.New("url和请求类型已存在")
    }
    result := l.svcCtx.Gorm.Model(&model.Permission{}).Where("id = ?", req.Id).Updates(up)
    if result.Error != nil {
        return 500000, nil, errors.New("更新权限失败")
    }
    upId := dataFormat.IntToString(req.Id)
    return 200000, &upId, nil
}

func (l *PermissionLogic) Index(req *types.PermissionSearchReq) (code int, resp []*types.PermissionAddReply, err error) {
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
    if req.Page >= 1 {
        err = validate.Var(req.Page, "number,lte=10000,gte=1")
    }
    if req.Limit >= 1 {
        err = validate.Var(req.Limit, "number,lte=1000,gte=1")
    }
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return 400000, nil, errors.New(dataFormat.RemoveTopStruct(transStr))
    }
    var getData model.Permission
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
    if req.Limit <= 0 {
        req.Limit = 10
    }
    if req.Page <= 0 {
        req.Page = 1
    }
    var all []*model.Permission
    var total int64
    db := l.svcCtx.Gorm.Where(&getData)
    if req.ParentId >= 0 {
        db = db.Where("parent_id =?", req.ParentId)
    }
    if req.Status >= 0 {
        db = db.Where("status =?", req.Status)
    }
    db.Count(&total)
    pageSetNum, offset := dataFormat.Page(req.Limit, req.Page, total)
    result := db.Limit(pageSetNum).Offset(offset).Find(&all)
    if result.Error != nil {
        return 500000, nil, errors.New("查询用户列表失败")
    }
    getAll := []*types.PermissionAddReply{}
    if len(all) <= 0 {
        return 200000, getAll, nil
    }
    for _, v := range all {
        r := &types.PermissionAddReply{
            Id:         int(v.Id),
            ParentId:   v.ParentId,
            ApiUrl:     v.ApiUrl,
            Method:     v.Method,
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
