package str

import (
	"math/rand"
	"slices"
	"strings"
)

const (
	upperLetters string = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	lowerLetters string = "abcdefghjkmnpqrstuvwxyz"
	numbers      string = "23456789"
)

// Create a random string of given length, with uppercase, lowercase letters, and numbers
func RandomString(length uint, useUpper, useLower, useNumber bool) string {
	charSource := ""
	if useUpper {
		charSource += upperLetters
	}
	if useLower {
		charSource += lowerLetters
	}
	if useNumber {
		charSource += numbers
	}
	numChars := len(charSource)
	if numChars == 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range length {
		idx := rand.Intn(numChars)
		b[i] = charSource[idx]
	}
	return string(b)
}

// Repeat string, joined by glue
func Repeat(repeat int, text, glue string) string {
	return strings.Join(slices.Repeat([]string{text}, repeat), glue)
}
