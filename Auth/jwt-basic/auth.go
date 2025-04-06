package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewUserClaims(email string, expiry time.Duration) UserClaims {
	return UserClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "myapp",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
}