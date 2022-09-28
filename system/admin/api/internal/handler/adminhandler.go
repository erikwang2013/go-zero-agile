package handler

import (
	"net/http"

	"erik-agile/common/errorx"
	"erik-agile/common/successx"
	"erik-agile/system/admin/api/internal/logic"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func adminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminInfoReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        code, resp, err := l.Admin(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code,resp))
        }
    }
}

func adminInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminInfoAllReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        code, resp, err := l.AdminInfo(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code,resp))
        }
    }
}

func adminCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminAddReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        code, resp, err := l.Create(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code,resp))
        }
    }
}

func adminDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.DeleteIdsReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        code, resp, err := l.Delete(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code,resp))
        }
    }
}

func adminPutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminPutReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        code, resp, err := l.Put(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code,resp))
        }
    }
}
