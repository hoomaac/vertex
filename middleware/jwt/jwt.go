package jwt

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

var secretKey = os.Getenv("JWT_SECRET")

func GenerateJwt() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foot": "bar",
		"nf":   "hey",
	})


	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatalf("generating token failed, %v", err)
		return ""
	}

	return tokenString
}
