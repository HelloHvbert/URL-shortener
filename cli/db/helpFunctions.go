package db

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

// If user does not provide a custom shortened URL, generate a random one
func GenerateShort() string {
	// Get the max length of the short URL
	length, err := strconv.ParseInt(os.Getenv("SHORT_URL_MAX_LEN"), 10, 64)
	if err != nil {
		length = 6
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano()) // Ensure different output on each run

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	// Check if the generated short URL already exists in the database
	// If it does, generate a new one
	isUsed, _ := GetOriginalUrl(string(result))

	for isUsed != "" {
		result = make([]byte, length)
		for i := range result {
			result[i] = charset[rand.Intn(len(charset))]
		}
		isUsed, _ = GetOriginalUrl(string(result))
	}
	
	return string(result)
}
