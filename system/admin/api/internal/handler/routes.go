// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"erik-agile/system/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/system/admin/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AdminMiddle},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/system/admin",
					Handler: adminHandler(serverCtx),
				},
                {
					Method:  http.MethodPost,
					Path:    "/system/admin",
					Handler: createHandler(serverCtx),
				},
                 {
					Method:  http.MethodDelete,
					Path:    "/system/admin",
					Handler: deleteHandler(serverCtx),
				},
                {
					Method:  http.MethodPut,
					Path:    "/system/admin",
					Handler: putHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
