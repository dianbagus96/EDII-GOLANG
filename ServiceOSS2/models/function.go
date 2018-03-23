package models

import (
	"math/rand"
	"time"
)

func shuffle(Title string) string {
	src := []string{
		"A", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	}

	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}

	var char string
	for p, _ := range final {
		char += final[p]
	}

	runes := []rune(char)
	var count int = 15 - len(Title)
	safeSubstring := Title + string(runes[0:count])
	return safeSubstring
}

func generateCode(Title string) string {
	shuffled := shuffle(Title)
	return shuffled
}
