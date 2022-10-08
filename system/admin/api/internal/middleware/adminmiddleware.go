package middleware

import (
	"erik-agile/common/errorx"
	commonData "erik-agile/system/admin/api/internal/common-data"
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/svc/gorm"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminMiddleware struct {
    config config.Config
    db     *gorm.Gormdb
}

func NewAdminMiddleware(gorm *gorm.Gormdb, c config.Config) *AdminMiddleware {
    return &AdminMiddleware{db: gorm, config: c}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        getToken := r.Header.Get("Authorization")
        if len(getToken) <= 0 {
            httpx.Error(w, errorx.NewCodeError(401000, "令牌认证失败"))
            return
        }
        result := commonData.CheckPermission(m.db.Gorm, r.Context(), m.config, r.RequestURI, r.Method)
        if false == result {
            httpx.Error(w, errorx.NewCodeError(403000, "非法授权"))
            return
        }
        next(w, r)
    }
}
