// Code generated by goctl. DO NOT EDIT!
// Source: admin.proto

package server

import (
	"context"

	"go-zero-agile/system/admin/rpc/internal/logic"
	"go-zero-agile/system/admin/rpc/internal/svc"
	"go-zero-agile/system/admin/rpc/types/admin"
)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminServer) GetAdmin(ctx context.Context, in *admin.IdRequest) (*admin.AdminResponse, error) {
	l := logic.NewGetAdminLogic(ctx, s.svcCtx)
	return l.GetAdmin(in)
}
