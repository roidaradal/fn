// Package str contains string-related functions.
package str

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var printer = message.NewPrinter(language.English)

// Utility function for checking empty string
func IsEmpty(text string) bool {
	return text == ""
}

// Utility function for checking non-empty string
func NotEmpty(text string) bool {
	return text != ""
}

// Return the number formatted with commas
func Comma[T ~int | ~uint](number T) string {
	return printer.Sprintf("%d", number)
}

// Split the text by separator, trim each part's extra whitespace
func CleanSplit(text, sep string) []string {
	parts := strings.Split(text, sep)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

// Split text by whitespace
func SpaceSplit(text string) []string {
	return strings.Fields(strings.TrimSpace(text))
}

// Split text by comma
func CommaSplit(text string) []string {
	return CleanSplit(text, ",")
}

// Get lines from text, separated by \n
func Lines(text string) []string {
	return CleanSplit(text, "\n")
}

// Join strings parts by glue
func Join(glue string, parts ...string) string {
	return strings.Join(parts, glue)
}

// Get the initial letters when text is split by separator
func PartInitials(text, sep string) string {
	parts := CleanSplit(text, sep)
	initials := make([]byte, len(parts))
	for i, part := range parts {
		initials[i] = part[0]
	}
	return string(initials)
}
