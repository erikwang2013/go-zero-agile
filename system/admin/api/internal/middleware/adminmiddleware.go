package middleware

import (
	"erik-agile/system/admin/api/internal/config"
	"net/http"
)

type AdminMiddleware struct {
    config config.Config
}

func NewAdminMiddleware(c config.Config) *AdminMiddleware {
    return &AdminMiddleware{config: c}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        next(w, r)
    }
}
