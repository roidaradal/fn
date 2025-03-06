package hash

import (
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

const (
	hashCost     int    = 10
	upperLetters string = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	lowerLetters string = "abcdefghjkmnpqrstuvwxyz"
	numbers      string = "23456789"
)

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	return string(bytes), err
}

func MatchPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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
