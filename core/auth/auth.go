package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"taskTracker/constants"
	"time"
)

var secretKey = []byte(constants.JwtKey) // todo fetch from config

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), // todo fetch Token expiration time from config
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
