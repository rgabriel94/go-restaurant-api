package testutil

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixMilli())
}

func RandomFloat() float64 {
	return rand.Float64() * 10
}

func RandomString(length int) string {
	var sb strings.Builder
	alphabetLen := len(alphabet)
	for i := 0; i < length; i++ {
		sb.WriteByte(alphabet[rand.Intn(alphabetLen)])
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(12)
}
