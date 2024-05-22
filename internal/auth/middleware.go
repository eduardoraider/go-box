package auth

import (
	"context"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		claims, err, code := validate(tokenString)
		if err != nil {
			http.Error(rw, err.Error(), code)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "user_name", claims.UserName)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
