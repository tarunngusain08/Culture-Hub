package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var JwtSecret = []byte(os.Getenv("JWTSecret"))

func GenerateToken(userID uint, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userID,
		"username": name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
