package random_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/virzz/utils/random"
)

func BenchmarkCodePow10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		random.Code(1)
		random.Code(2)
		random.Code(3)
		random.Code(4)
		random.Code(5)
		random.Code(6)
		random.Code(7)
		random.Code(8)
		random.Code(9)
	}
}
func randomCode(n int) string {
	switch n {
	case 1:
		return strconv.Itoa(rand.Intn(9) + 1)
	case 2:
		return strconv.Itoa(rand.Intn(90) + 10)
	case 3:
		return strconv.Itoa(rand.Intn(900) + 100)
	case 4:
		return strconv.Itoa(rand.Intn(9000) + 1000)
	case 5:
		return strconv.Itoa(rand.Intn(90000) + 10000)
	case 6:
		return strconv.Itoa(rand.Intn(900000) + 100000)
	case 7:
		return strconv.Itoa(rand.Intn(9000000) + 1000000)
	case 8:
		return strconv.Itoa(rand.Intn(90000000) + 10000000)
	case 9:
		return strconv.Itoa(rand.Intn(900000000) + 100000000)
	case 10:
		return strconv.Itoa(rand.Intn(9000000000) + 1000000000)
	}
	return ""
}
func BenchmarkCodeSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomCode(1)
		randomCode(2)
		randomCode(3)
		randomCode(4)
		randomCode(5)
		randomCode(6)
		randomCode(7)
		randomCode(8)
		randomCode(9)
		randomCode(10)
	}
}
