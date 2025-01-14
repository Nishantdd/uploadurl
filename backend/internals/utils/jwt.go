package utils

import (
	"time"

	"github.com/Nishantdd/uploadurl/backend/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID uint) (string, error) {
	secret := []byte(config.Load().JWT.JWTSecret)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 48).Unix(), // Token expiration (48 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
