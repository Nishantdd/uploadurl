package utils

import (
	cryptRand "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	mathRand "math/rand/v2"
	"strings"
)

func Hash(data ...string) string {
	var input strings.Builder
	for _, s := range data {
		input.WriteString(s)
	}

	h := sha256.New()
	h.Write([]byte(input.String()))
	hashedBytes := h.Sum(nil)
	return fmt.Sprintf("%x", hashedBytes)
}

func CompareHash(hashedString string, data ...string) error {
	inputHash := Hash(data...)
	if inputHash == hashedString {
		return nil
	}
	return errors.New("hashes did not match")
}

func GenerateState() string {
	b := make([]byte, 32)
	cryptRand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

// Takes length as input and returns a alphanumeric string of given length
func GenerateUniqueString(len int) string {
	var slug strings.Builder

	allowedCharacters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12334567890"
	for i := 0; i < len; i++ {
		slug.WriteByte(allowedCharacters[mathRand.IntN(16)])
	}

	return slug.String()
}
