package util

import (
	"math/rand"
	"strings"
	"time"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

// this init function is autically called when the package is initialized
func init() {

	//get the current time in nanao seconds as an integer value
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random number between min & max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random string of Owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney to generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EUR"}
	n := len(currencies)
	return currencies[rand.Intn(n)] //generate a random integer between 0 (inclusive) and n

}
