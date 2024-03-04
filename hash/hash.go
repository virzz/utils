package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/virzz/utils/random"
)

func Sha256(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}

func HmacSha256(salt string) string {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write(random.Bytes(16))
	return hex.EncodeToString(h.Sum(nil))
}
