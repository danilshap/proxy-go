package utils

import (
	"crypto/rand"
	"fmt"
)

const baseUrlDomain string = "localhost:8081"

func GenerateUri() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := make([]byte, 10)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	route := string(bytes)

	return fmt.Sprintf("http://%s/%s", baseUrlDomain, route)
}
