package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(uid uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": 	uid,
		"exp":		time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(t string) (uint, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return uint(claims["user_id"].(float64)), nil
}
