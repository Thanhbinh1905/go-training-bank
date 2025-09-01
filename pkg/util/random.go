package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates random int between min and max
func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomID() int64 {
	return randomInt(1, 1000)
}

func RandomMoney() int64 {
	return randomInt(0, 1000)
}

func RandomEntry() int64 {
	return randomInt(-50, 50)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, VND}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomEmail(username string) string {
	return fmt.Sprintf("%s@gmail.com", username)
}
