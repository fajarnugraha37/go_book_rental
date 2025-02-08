package helper

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
)

func GenerateRandomBase32(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return base32.StdEncoding.EncodeToString(randomBytes)[:length], nil // Take first 'length' chars
}
