package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		claims := new(Claims)

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(rw, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "user_name", claims.UserName)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
