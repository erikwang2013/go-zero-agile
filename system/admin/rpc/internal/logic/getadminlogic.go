package logic

import (
	"context"

	"erik-agile/system/admin/rpc/internal/svc"
	"erik-agile/system/admin/rpc/types/admin"

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
    one,err:=l.svcCtx.AdminModel.FindOne(l.ctx, in.Id)
    if err != nil {
        return nil, err
    }
	return &admin.AdminResponse{
        Id:one.Id,
        Name: one.Name,
        Gender: one.Gender,
        Phone: one.Phone,
        Email: one.Email,
        Status: one.Status,
        IsDelete: one.IsDelete,
        CreateTime: one.CreateTime.Format("006-01-02 15:04:05"),
        UpdateTime: one.UpdateTime.Format("006-01-02 15:04:05"),
    }, nil
}
