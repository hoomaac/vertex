package jwt

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = os.Getenv("JWT_SECRET")

func GenerateJwt(username string, email string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
		"IssuedAt": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatalf("generating token failed, %v", err)
		return ""
	}

	return tokenString
}
