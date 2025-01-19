package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
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
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
