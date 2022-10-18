package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func RandomMoney() int64 {
	return int64(RandomInt(1, 100000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	return currencies[rand.Intn(len(currencies))]
}
