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
                    Method:  http.MethodPost,
                    Path:    "/system/admin",
                    Handler: adminCreateHandler(serverCtx),
                },
                {
                    Method:  http.MethodDelete,
                    Path:    "/system/admin",
                    Handler: adminDeleteHandler(serverCtx),
                },
                {
                    Method:  http.MethodPut,
                    Path:    "/system/admin",
                    Handler: adminPutHandler(serverCtx),
                },
                {
                    Method:  http.MethodGet,
                    Path:    "/system/admin",
                    Handler: adminHandler(serverCtx),
                },
                 {
                    Method:  http.MethodGet,
                    Path:    "/system/admin/info",
                    Handler: adminInfoHandler(serverCtx),
                },
                {
                    Method:  http.MethodPost,
                    Path:    "/system/permission",
                    Handler: permissionCreateHandler(serverCtx),
                },
                {
                    Method:  http.MethodDelete,
                    Path:    "/system/permission",
                    Handler: permissionDeleteHandler(serverCtx),
                },
                {
                    Method:  http.MethodPut,
                    Path:    "/system/permission",
                    Handler: permissionPutHandler(serverCtx),
                },
                {
                    Method:  http.MethodGet,
                    Path:    "/system/permission",
                    Handler: permissionHandler(serverCtx),
                },
                 {
                    Method:  http.MethodPost,
                    Path:    "/system/role",
                    Handler: roleCreateHandler(serverCtx),
                },
                {
                    Method:  http.MethodDelete,
                    Path:    "/system/role",
                    Handler: roleDeleteHandler(serverCtx),
                },
                {
                    Method:  http.MethodPut,
                    Path:    "/system/role",
                    Handler: rolePutHandler(serverCtx),
                },
                {
                    Method:  http.MethodGet,
                    Path:    "/system/role",
                    Handler: roleHandler(serverCtx),
                },
            }...,
        ),
        rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
    )
}
