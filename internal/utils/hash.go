package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomID generates a random ID as a hexadecimal string
func GenerateRandomID() string {
	// Define the length of the ID in bytes (e.g., 16 bytes for 128-bit ID)
	length := 16
	// Create a byte slice to hold the random data
	randomBytes := make([]byte, length)
	// Generate random bytes
	rand.Read(randomBytes)
	// Encode the random bytes as a hexadecimal string
	userID := hex.EncodeToString(randomBytes)

	return userID
}
