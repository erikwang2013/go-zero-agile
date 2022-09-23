package logic

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"time"

	"erik-agile/common/data"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"
	"erik-agile/system/admin/model"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
    return &LoginLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

var trans ut.Translator

func validateRegister(v *validator.Validate) {
    v.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return "j"
        }
        return name
    })
    zh := zh.New()
    uni := ut.New(zh)
    trans, _ = uni.GetTranslator("zh")
    zh_translations.RegisterDefaultTranslations(v, trans)
    return
}

func (l *LoginLogic) Login(req *types.LoginReq) (reqly *types.LoginReply, err error) {
    validate := validator.New()
    validateRegister(validate)
    err = validate.Struct(req)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        transStr := varError.Translate(trans)
        return nil, errors.New(data.RemoveTopStruct(transStr))
    }
    adminInfo, err := l.svcCtx.AdminModel.FindOneName(l.ctx, req.UserName)
    if err != nil {
        return nil, errors.New("登录校验异常")
    }
    if strings.Compare(adminInfo.Password, req.Password) != 0 {
        return nil, errors.New("用户名或密码错误")
    }
    token, now, accessExpire, err := l.getJwtToken(adminInfo.Id)
    if err != nil {
        return nil, errors.New("令牌生成失败")
    }
    getTime := time.Unix(time.Now().Unix(), 0)
    adminLog := &model.AdminLoginLog{
        Id:          data.NextSonyFlakeIdInt64(),
        AdminId:     adminInfo.Id,
        AccessToken: token,
        LoginTime:   getTime,
    }
    go l.svcCtx.AdminLoginLogModel.Insert(l.ctx, adminLog)
    return &types.LoginReply{
        Id:           adminInfo.Id,
        Name:         adminInfo.Name,
        AccessToken:  token,
        AccessExpire: now + accessExpire,
        RefreshAfter: now + accessExpire/2,
    }, nil
}

func (l *LoginLogic) getJwtToken(adminId int) (string, int64, int64, error) {
    iat := time.Now().Unix()
    secretKey := l.svcCtx.Config.Auth.AccessSecret
    seconds := l.svcCtx.Config.Auth.AccessExpire
    claims := make(jwt.MapClaims)
    claims["exp"] = iat + seconds
    claims["iat"] = iat
    claims["admin_id"] = adminId
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims = claims
    accessToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        logx.Error(err)
        return "", iat, seconds, err
    }
    return accessToken, iat, seconds, nil
}
