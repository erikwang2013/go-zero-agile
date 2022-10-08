package middleware

import (
	"erik-agile/common/errorx"
	commonData "erik-agile/system/admin/api/internal/common-data"
	"erik-agile/system/admin/api/internal/svc/gorm"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminMiddleware struct {
    db *gorm.Gormdb
}

func NewAdminMiddleware(gorm *gorm.Gormdb) *AdminMiddleware {
    return &AdminMiddleware{db: gorm}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        getToken := r.Header.Get("Authorization")
        if len(getToken) <= 0 {
            httpx.Error(w, errorx.NewCodeError(401000, "令牌认证失败"))
            return
        }
        // getId := commonData.GetAdminId(r.Context())
        // logx.Error("===获取id==")
        // logx.Error(getId)
        result := commonData.CheckPermission(m.db.Gorm, r.Context(), r.RequestURI, r.Method)
        if false == result {
            httpx.Error(w, errorx.NewCodeError(403000, "非法授权"))
            return
        }
        next(w, r)
    }
}
