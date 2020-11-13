package auth

import (
	"context"
	"github.com/golang_project_01_server/services_authorization"
	"net/http"
	//"fmt"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "X-Authorization", services_authorization.AuthToken.Token)

			k := ctx.Value("X-Authorization")

			if k != nil {
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		})
	}
}