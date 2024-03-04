package random

import (
	"encoding/hex"
	"math"
	"math/rand"
	"strconv"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func ByteUniq(n int) map[byte]struct{} {
	sm := map[byte]struct{}{}
	for len(sm) < n {
		sm[letterBytes[rand.Int63()%int64(len(letterBytes))]] = struct{}{}
	}
	return sm
}

func Bytes(n int) []byte {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func BytesHex(n int) string {
	return hex.EncodeToString(Bytes(n))
}

func String(n int) string {
	return string(Bytes(n))
}

func Code(bit ...int) string {
	_bit := 6
	if len(bit) > 0 {
		_bit = bit[0]
	}
	x := int(math.Pow10(_bit - 1))
	return strconv.Itoa(rand.Intn(int(math.Pow10(_bit))-x) + x)
}
