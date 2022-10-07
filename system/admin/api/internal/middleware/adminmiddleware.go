package middleware

import (
	"context"
	dataFormat "erik-agile/common/data-format"
	"erik-agile/common/errorx"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
    return &AdminMiddleware{}
}

//获取用户id
func GetAdminId(ctx context.Context) int {
    adminId := ctx.Value("admin_id")
    getAdminId := fmt.Sprintf("%v", adminId)
    return dataFormat.StringToInt(getAdminId)
}

//校验权限
func CheckPermission(checkStr string) bool {
    return true
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        getToken := r.Header.Get("Authorization")
        if len(getToken) <= 0 {
            httpx.Error(w, errorx.NewCodeError(401000, "令牌认证失败"))
            return
        }
        checkStr := dataFormat.GetMd5(r.RequestURI + r.Method)
        result := CheckPermission(checkStr)
        if false == result {
            httpx.Error(w, errorx.NewCodeError(403000, "非法授权"))
            return
        }
        next(w, r)
    }
}
