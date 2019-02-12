package random

import (
	"math/rand"
	"time"
)

var randomSource = rand.NewSource(time.Now().UnixNano())

const (
	lettersBytes  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func RandomString(n int) string {
	b := make([]byte, n)
	cache, remain := randomSource.Int63(), letterIdxMax
	for i := n - 1; i >= 0; {
		if remain == 0 {
			cache, remain = randomSource.Int63(), letterIdxMax
		}
		idx := int(cache & letterIdxMask)
		if idx < len(lettersBytes) {
			b[i] = lettersBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
