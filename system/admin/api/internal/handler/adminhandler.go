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

func adminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        logx.Info("====打印handler=1===")
        logx.Info(r)
        var req types.AdminInfoReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        logx.Info("====打印handler=2===")
        logx.Info(req)
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        resp, err := l.Admin(&req)
        if err != nil {
            httpx.Error(w, errorx.NewDefaultError(err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(resp))
        }
    }
}

func createHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminAddReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        resp, err := l.Create(&req)
        if err != nil {
            httpx.Error(w, errorx.NewDefaultError(err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(resp))
        }
    }
}

func deleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminInfoReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        resp, err := l.Delete(&req)
        if err != nil {
            httpx.Error(w, errorx.NewDefaultError(err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(resp))
        }
    }
}

func putHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.AdminInfoReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
        l := logic.NewAdminLogic(r.Context(), svcCtx)
        resp, err := l.Put(&req)
        if err != nil {
            httpx.Error(w, errorx.NewDefaultError(err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(resp))
        }
    }
}
