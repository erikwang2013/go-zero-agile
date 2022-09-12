package logic

import (
	"context"

	"go-zero-agile/system/admin/rpc/internal/svc"
	"go-zero-agile/system/admin/rpc/types/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminLogic {
	return &GetAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminLogic) GetAdmin(in *admin.IdRequest) (*admin.AdminResponse, error) {
	// todo: add your logic here and delete this line

	return &admin.AdminResponse{}, nil
}
