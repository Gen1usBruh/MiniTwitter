package random

import (
	"math/rand"
	"time"
	"math"
)

func GenerateRandomAlphanumeric(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRandomInt32() int32 {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return int32(seededRand.Intn(math.MaxInt32))
}