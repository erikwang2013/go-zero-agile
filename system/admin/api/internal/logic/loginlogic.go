package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"

	validator "github.com/go-playground/validator/v10"
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

type loginData struct {
    Username string `json:"username" validate:"gte>=4,lte<=20"`
    Password string `json:"password" validate:"gte>10,lte<=200"`
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
    val := validator.New()
    loginData := &loginData{
        Username: req.Username,
        Password: req.Password,
    }
    err = val.Struct(loginData)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        return nil, varError
    }
    adminInfo, err := l.svcCtx.AdminModel.FindOneName(l.ctx, req.Username)
    if err != nil {
        return nil, err
    }
    if strings.Compare(adminInfo.Password, req.Password) != 0 {
        return nil, errors.New("用户名或密码错误")
    }
    token, now, accessExpire, err := l.getJwtToken(adminInfo.Id)
    if err != nil {
        return nil, err
    }
    return &types.LoginReply{
        Id:           adminInfo.Id,
        Name:         adminInfo.Name,
        AccessToken:  token,
        AccessExpire: now + accessExpire,
        RefreshAfter: now + accessExpire/2,
    }, nil
}

func (l *LoginLogic) getJwtToken(adminId uint) (string, uint64, uint64, error) {
    iat := uint64(time.Now().Unix())
    secretKey := l.svcCtx.Config.Auth.AccessSecret
    seconds := uint64(l.svcCtx.Config.Auth.AccessExpire)
    claims := make(jwt.MapClaims)
    claims["exp"] = iat + seconds
    claims["iat"] = iat
    claims["admin_id"] = adminId
    token := jwt.New(jwt.SigningMethodES256)
    token.Claims = claims
    accessToken, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", iat, seconds, err
    }
    return accessToken, iat, seconds, nil
}
