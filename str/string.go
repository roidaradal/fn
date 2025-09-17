package str

import (
	"strings"

	"github.com/roidaradal/fn"
)

func CleanSplit(text string, sep string) []string {
	parts := strings.Split(text, sep)
	parts = fn.Map(parts, strings.TrimSpace)
	return parts
}

func SpaceSplit(text string) []string {
	return strings.Fields(strings.TrimSpace(text))
}
