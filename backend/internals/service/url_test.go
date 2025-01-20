package service

import (
	"fmt"
	"testing"
)

// Throws error with provided message if the condition doesn't match
func assert(t *testing.T, condition bool, msg string) {
	if !condition {
		t.Error(msg)
	}
}

// Tests wheter the string has randomness to it or not
func isRandomString(s *string) bool {
	if len(*s) == 0 {
		return false
	}
	for i := 1; i < len(*s); i++ {
		if (*s)[i] != (*s)[0] {
			return true
		}
	}
	return false
}

func TestUrlShortner(t *testing.T) {
	sizes := []int{8, 16, 32}
	for _, size := range sizes {
		shortUrl := GenerateRandomSlug(size)
		assert(t, len(shortUrl) == size, fmt.Sprintf("expected length %d, but got %d", size, len(shortUrl)))
		assert(t, isRandomString(&shortUrl), fmt.Sprintf("%s is missing randomness", shortUrl))
	}
}
