package controllers

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"email"`
	jwt.StandardClaims
}

var Secret = []byte("secret")

func createTokenJWT(name, email string) (string, error) {
	claims := &jwtCustomClaims{
		name,
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(Secret)
	if err != nil {
		log.Println("Error: ", err)
	}

	return t, err
}
