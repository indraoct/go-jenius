package jenius

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func Base64Encode(s string) (string) {
	data := []byte(s)
	encode := base64.StdEncoding.EncodeToString(data)

	return encode
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}