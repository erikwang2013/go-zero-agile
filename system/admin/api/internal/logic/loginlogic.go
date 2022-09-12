package logic

import (
	"context"

	"go-zero-agile/system/admin/api/internal/svc"
	"go-zero-agile/system/admin/api/internal/types"

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

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line

	return
}

func (l *LoginLogic) getJwtToken(secretKey string, iat,seconds,adminId uint64) (string,error){
    claims:=make(jwt.MapClaims)
    claims["exp"]=iat+seconds
    claims["iat"]=iat
    claims["adminId"]=adminId
    token:=jwt.New(jwt.SigningMethodES256)
    token.Claims=claims
    return token.SignedString([]byte(secretKey))
}