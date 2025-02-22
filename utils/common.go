package utils


import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	byteLength := (length * 6) / 8 
	if (length*6)%8 != 0 {
		byteLength++
	}

	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	return randomString[:length], nil
}