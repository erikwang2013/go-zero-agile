package middleware

import (
	"erik-agile/common/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
    return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        accessToken := r.Header.Get("Authorization")
        if len(accessToken) <= 5 {
            httpx.Error(w, errorx.NewDefaultError("认证令牌错误或非法"))
        }
        next(w, r)
    }
}
