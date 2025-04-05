package fn

import (
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ParseInt(value string) int {
	number, err := strconv.Atoi(strings.TrimSpace(value))
	if err == nil {
		return number
	} else {
		return 0
	}
}

func ParseFloat(value string) float64 {
	number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if err == nil {
		return number
	} else {
		return 0
	}
}
