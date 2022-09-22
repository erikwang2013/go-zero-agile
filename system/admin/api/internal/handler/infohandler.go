package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"erik-agile/system/admin/api/internal/logic"
	"erik-agile/system/admin/api/internal/svc"
	"erik-agile/system/admin/api/internal/types"
)

func infoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
