package middleware

import (
	"gcloud/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		// 注入用户信息到上下文
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserName", uc.Name)
		r.Header.Set("UserIdentity", uc.Identity)
		// Passthrough to next handler if need
		next(w, r)
	}
}
