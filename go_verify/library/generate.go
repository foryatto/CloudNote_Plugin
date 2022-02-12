package library

import (
	"math/rand"
	"strings"
	"time"
)

const NUMBERS = "0123456789"

func NewRandomNums(len int) string {
	if len <= 0 || len >= 10 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	var result strings.Builder
	for i := 0; i < len; i++ {
		result.WriteByte(NUMBERS[rand.Intn(10)])
	}
	return result.String()
}
