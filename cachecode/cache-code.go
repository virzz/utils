package cachecode

import (
	"strings"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var (
	cacheCodeMap = cmap.New[string]()
)

func CheckCode(key, code string) bool {
	v, ok := cacheCodeMap.Pop(key)
	return ok && (v == code || v == strings.ToLower(code))
}

func CacheCode(key, code string) (err error) {
	cacheCodeMap.Set(key, strings.ToLower(code))
	go func() {
		<-time.After(10 * time.Minute)
		cacheCodeMap.Pop(key)
	}()
	return nil
}
