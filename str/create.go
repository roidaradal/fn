package str

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

const (
	upperLetters string = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	lowerLetters string = "abcdefghjkmnpqrstuvwxyz"
	numbers      string = "23456789"
)

// Create random string of given length, using uppercase, lowercase letters, and numbers (flags)
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
		idx := rand.IntN(numChars)
		b[i] = charSource[idx]
	}
	return string(b)
}

// Repeat string, joined by glue
func Repeat(repeat int, text, glue string) string {
	return strings.Join(slices.Repeat([]string{text}, repeat), glue)
}

// Wrap string in backticks
func WrapBackticks(text string) string {
	return fmt.Sprintf("`%s`", text)
}

// Create string of items separated by comma, wrapped in curly braces
func WrapBraces[T any](items []T) string {
	return "{ " + strings.Join(List(items), ", ") + " }"
}

// Create string of items separated by comma, wrapped in square brackets
func WrapBrackets[T any](items []T) string {
	return "[ " + strings.Join(List(items), ", ") + " ]"
}

// Create string of items separated by comma, wrapped in parentheses
func WrapParens[T any](items []T) string {
	return "( " + strings.Join(List(items), ", ") + " )"
}

type Builder struct {
	items []string
}

func NewBuilder() *Builder {
	return &Builder{
		items: make([]string, 0),
	}
}

func (b *Builder) Add(item string) {
	b.items = append(b.items, item)
}

func (b *Builder) Build(separator string) string {
	return strings.Join(b.items, separator)
}
