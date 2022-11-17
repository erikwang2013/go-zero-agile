package middleware

import (
	"erik-agile/system/admin/api/internal/config"
	"erik-agile/system/admin/api/internal/svc/gorm"
	"net/http"
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
        next(w, r)
    }
}
