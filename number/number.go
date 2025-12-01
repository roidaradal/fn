// Package number contains number-related functions.
package number

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

// Number interface unifies the number types
type Number interface {
	~int | ~uint | ~float32 | ~float64
}

// Absolute value of integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Compute ratio of numerator / denominator ints
func Ratio[T Number](numerator, denominator T) float64 {
	return float64(numerator) / float64(denominator)
}

// Parse integer value, defaults to 0 if invalid
func ParseInt(value string) int {
	number, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0
	}
	return number
}

// Parse float64 value, defaults to 0 if invalid
func ParseFloat(value string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err != nil {
		return 0
	}
	return number
}

// Convert binary number string to integer, default to 0 if invalid
func ParseBinary(value string) int {
	number, err := strconv.ParseInt(strings.TrimSpace(value), 2, 64)
	if err != nil {
		return 0
	}
	return int(number)
}

// Return the number formatted with commas
func Comma[T ~int | ~uint](number T) string {
	return printer.Sprintf("%d", number)
}

// Return the float number formatted with commas, up to given number of decimals
func FloatComma[T ~float32 | ~float64](number T, decimals int) string {
	template := fmt.Sprintf("%%.%df", decimals)
	return printer.Sprintf(template, number)
}
