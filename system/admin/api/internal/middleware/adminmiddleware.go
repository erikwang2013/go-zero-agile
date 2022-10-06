package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
    return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
        logx.Error("打印中间件")
    return func(w http.ResponseWriter, r *http.Request) {
    logx.Error("打印url")
    //getUrl:=r.RequestURI
    //getMethod:=r.Method
        next(w, r)
    }
}
