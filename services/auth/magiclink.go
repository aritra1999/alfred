package auth

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateMagicLinkToken() string {
	tokenBytes := make([]byte, 32)
	rand.Read(tokenBytes)
	return base64.StdEncoding.EncodeToString(tokenBytes)
}
