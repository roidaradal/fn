package conv

import (
	"strconv"
	"strings"

	"github.com/roidaradal/fn"
)

// Absolute value of integer
func Abs(x int) int {
	return fn.Ternary(x < 0, -x, x)
}

// Parse integer value, default to 0 if invalid
func ParseInt(value string) int {
	number, err := strconv.Atoi(strings.TrimSpace(value))
	return fn.Ternary(err == nil, number, 0)
}

// Parse float64 value, default to 0 if invalid
func ParseFloat(value string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	return fn.Ternary(err == nil, number, 0)
}

// Convert binary number string to integer, default to 0 if invalid
func ParseBinary(value string) int {
	number, err := strconv.ParseInt(strings.TrimSpace(value), 2, 64)
	return fn.Ternary(err == nil, int(number), 0)
}
