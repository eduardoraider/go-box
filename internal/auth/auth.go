package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func GetSecret() string {
	return jwtSecret
}

type Claims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func createToken(authenticated Authenticated) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		UserID:   authenticated.GetID(),
		UserName: authenticated.GetName(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(GetSecret()))
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authenticated interface {
	GetID() int64
	GetName() string
}

type authenticateFunc func(string, string) (Authenticated, error)

type handler struct {
	authenticate authenticateFunc
}

func (h *handler) auth(creds Credentials) (token string, err error, code int) {
	u, err := h.authenticate(creds.Username, creds.Password)
	if err != nil {
		return "", err, http.StatusUnauthorized
	}

	token, err = createToken(u)
	if err != nil {
		return "", err, http.StatusInternalServerError
	}

	return token, nil, http.StatusOK
}
