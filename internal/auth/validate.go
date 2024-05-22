package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func validate(tokenString string) (*Claims, error, int) {
	claims := new(Claims)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSecret()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, err, http.StatusUnauthorized
		}
		return nil, err, http.StatusBadRequest
	}

	if !token.Valid {
		return nil, err, http.StatusUnauthorized
	}

	return claims, nil, http.StatusOK
}
