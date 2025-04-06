package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey 	string
	expiry		time.Duration
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

func NewJWTManager(secretKey string, expiry time.Duration) *JWTManager {
	return &JWTManager{
		secretKey: 	secretKey,
		expiry: 	expiry,
	}
}

// Generates token create a signed token with user claims.
func (j *JWTManager) GenerateToken(email string) (string, error) {
	claims := NewUserClaims(email, j.expiry)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) ValidateToken(token string) (*UserClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &UserClaims{}, func (parsedToken *jwt.Token) (interface{}, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		} 
		return nil, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(*UserClaims)
	if !ok || !parsedToken.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}