package handler

import (
	"erik-agile/common/errorx"
	"erik-agile/common/successx"
	"erik-agile/system/admin/api/internal/logic"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func roleCreateHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.RoleAddReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewRoleLogic(r.Context(), svcCtx)
        code, resp, err := l.Create(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}

func roleDeleteHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.DeleteIdsReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewRoleLogic(r.Context(), svcCtx)
        code, resp, err := l.Delete(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}

func rolePutHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.RolePutReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewRoleLogic(r.Context(), svcCtx)
        code, resp, err := l.Put(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}

func roleHandler(svcCtx *svc.ServiceContext ) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.RoleSearchReq
        if err := httpx.Parse(r, &req); err != nil {
            logx.Error(err)
            httpx.Error(w, errorx.NewCodeError(401000, "请求参数错误"))
            return
        }
        l := logic.NewRoleLogic(r.Context(), svcCtx)
        code, resp, err := l.Index(&req)
        if err != nil {
            httpx.Error(w, errorx.NewCodeError(code, err.Error()))
        } else {
            httpx.OkJson(w, successx.NewDefaultSuccess(code, resp))
        }
    }
}
