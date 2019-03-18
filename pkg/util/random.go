package util

import (
	"math/rand"
	"time"
)

type Random struct {
	rand *rand.Rand
}

func NewRandom() *Random {
	return NewSeededRandom(time.Now().UTC().UnixNano())
}

func NewSeededRandom(seed int64) *Random {
	return &Random{rand: rand.New(rand.NewSource(seed))}
}

// return a random string of the given length from a standard lowercase-only alphanumeric alphabet)
func (r *Random) RandomString(length int) string {
	return r.RandomStringUsingCustomAlphabet(length, []rune("0123456789abcdefghijklmnopqrstuvwxyz"))
}

// return a random string of the given length from a custom alphabet.
func (r *Random) RandomStringUsingCustomAlphabet(length int, alphabet []rune) string {
	res := make([]rune, length)
	for i := range res {
		res[i] = alphabet[r.rand.Intn(len(alphabet))]
	}

	return string(res)
}
