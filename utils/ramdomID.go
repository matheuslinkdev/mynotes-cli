package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(charset))
		sb.WriteByte(charset[randomIndex])
	}

	return sb.String()
}