package middleware

import (
	"erik-agile/common/errorx"
	commonData "erik-agile/system/admin/api/internal/common-data"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
    return &AdminMiddleware{}
}


func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        getToken := r.Header.Get("Authorization")
        if len(getToken) <= 0 {
            httpx.Error(w, errorx.NewCodeError(401000, "令牌认证失败"))
            return
        }
        getId := r.Context().Value("admin_id")
        logx.Error("===获取id==")
        logx.Error(getId)
         result := commonData.CheckPermission(r.RequestURI,r.Method)
        if false == result {
            httpx.Error(w, errorx.NewCodeError(403000, "非法授权"))
            return
        }
        next(w, r)
    }
}
