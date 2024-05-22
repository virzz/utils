package hash

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha1(data string) string {
	sum := sha1.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

func HmacSha1(salt, data string) string {
	h := hmac.New(sha1.New, []byte(salt))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

func HmacSha256(salt, data string) string {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
