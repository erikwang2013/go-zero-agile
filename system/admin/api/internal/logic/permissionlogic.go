package logic

import (
	"context"

	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"

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

func (l *PermissionLogic) Create(req *types.AdminInfoReq) (code int, resp []*types.AdminInfoReply, err error) {
    return 200000, nil, nil
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