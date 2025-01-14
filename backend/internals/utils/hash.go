package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(input string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func CompareHash(hashedString, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(input))
}
