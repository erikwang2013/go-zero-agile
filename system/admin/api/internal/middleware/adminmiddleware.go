package middleware

import (
	"net/http"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
    return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        next(w, r)
    }
}
