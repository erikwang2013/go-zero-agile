package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"

	"github.com/go-playground/validator/v10"
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

type LoginData struct {
    UserName string `json:"username" validate:"alphanum"`
    PassWord string `json:"password" validate:"alphanum"`
}

func (l *LoginLogic) Login(req *types.LoginReq) (reqly *types.LoginReply, err error) {
    validate := validator.New()
    loginVal :=&LoginData{
        UserName: req.UserName,
        PassWord: req.PassWord,
    }
    err = validate.Struct(loginVal)
    if err != nil {
        varError := err.(validator.ValidationErrors)
        return nil, errors.New(varError.Error())
    }
    adminInfo, err := l.svcCtx.AdminModel.FindOneName(l.ctx, req.UserName)
    if err != nil {
        return nil, errors.New("登录校验异常")
    }
    if strings.Compare(adminInfo.Password, req.PassWord) != 0 {
        return nil, errors.New("用户名或密码错误")
    }
    token, now, accessExpire, err := l.getJwtToken(1)
    if err != nil {
        return nil, errors.New("令牌生成失败")
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
