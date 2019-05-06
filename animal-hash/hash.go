package animalhash

import (
	"math"
	"strings"
)

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// Hash returns a unique string of animal えもじ
func Hash(s string, prime int) string {
	var hashSlice []string

	lenS := len(s)
	numberAnimals := len(Animals)

	for i := 0; i < lenS; i++ {
		index := pow(prime, lenS-(i+1)) * int(s[i]) % numberAnimals
		emoji := Animals[index]
		hashSlice = append(hashSlice, emoji)
	}
	return strings.Join(hashSlice, "")
}
