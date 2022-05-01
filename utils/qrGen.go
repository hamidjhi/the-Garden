package utils

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateQr() string {
	rand.Seed(time.Now().Unix())
	charSet := []rune("abcdedfghijklmnopqrst£")
	var output strings.Builder
	length := 10
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}
	return output.String()
}