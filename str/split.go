package str

import (
	"strings"

	"github.com/roidaradal/fn"
)

// Split the text by separator, then trim each part's extra whitespace
func CleanSplit(text string, sep string) []string {
	parts := strings.Split(text, sep)
	parts = fn.Map(parts, strings.TrimSpace)
	return parts
}

// Split the text by whitespace
func SpaceSplit(text string) []string {
	return strings.Fields(strings.TrimSpace(text))
}

// Split the text by comma
func CommaSplit(text string) []string {
	return CleanSplit(text, ",")
}

// Join string parts by the glue
func Join(glue string, parts ...string) string {
	return strings.Join(parts, glue)
}

// Get the lines from text, separated by \n
func Lines(text string) []string {
	return CleanSplit(text, "\n")
}

// Get the initial letters when text is split by glue
func SplitInitials(text string, glue string) string {
	initials := fn.Map(CleanSplit(text, glue), func(part string) byte {
		return part[0]
	})
	return string(initials)
}
