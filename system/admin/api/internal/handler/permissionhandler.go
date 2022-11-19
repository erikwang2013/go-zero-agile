package handler

import (
	"net/http"

	"erik-agile/common/errorx"
	"erik-agile/common/successx"
	"erik-agile/system/admin/api/internal/logic"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func permissionCreateHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.PermissionAddReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewPermissionLogic(r.Context(), svcCtx)
        code, resp, err := l.Create(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}

func permissionDeleteHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.DeleteIdsReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewPermissionLogic(r.Context(), svcCtx)
        code, resp, err := l.Delete(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}

func permissionPutHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.PermissionPutReq
        if err := httpx.Parse(r, &req); err != nil {
           logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewPermissionLogic(r.Context(), svcCtx)
        code, resp, err := l.Put(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}

func permissionHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.PermissionSearchReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewPermissionLogic(r.Context(), svcCtx)
        code, resp, err := l.Index(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}
