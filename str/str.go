// Package str contains string-related functions.
package str

import (
	"fmt"
	"strings"
	"unicode"
)

// Utility for returning string length
func Length(text string) int {
	return len(text)
}

// Utility function for checking empty string
func IsEmpty(text string) bool {
	return text == ""
}

// Utility function for checking non-empty string
func NotEmpty(text string) bool {
	return text != ""
}

// Split the text by separator, trim each part's extra whitespace
func CleanSplit(text, sep string) []string {
	parts := strings.Split(text, sep)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}

// Split the text by separator, maximum of N parts, trim each part's extra whitespace
func CleanSplitN(text, sep string, count int) []string {
	parts := strings.SplitN(text, sep, count)
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

// Get lines from text, separated by \n,
// Each line is trimmed for whitespace at both ends
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

// Center the string by padding with whitespace
func Center(text string, width int) string {
	padTotal := width - len(text)
	if padTotal <= 0 {
		return text // return as-is if no need to pad
	}
	pad1 := padTotal / 2
	pad2 := padTotal - pad1
	lpad, rpad := strings.Repeat(" ", pad2), strings.Repeat(" ", pad1)
	return lpad + text + rpad
}

// Left align given string
func LeftAlign(text string, width int) string {
	template := fmt.Sprintf("%%-%ds", width)
	return fmt.Sprintf(template, text)
}

// Right align given string
func RightAlign(text string, width int) string {
	template := fmt.Sprintf("%%%ds", width)
	return fmt.Sprintf(template, text)
}

// Check if string starts with uppercase letter
func StartsWithUpper(text string) bool {
	first := text[0]
	return 'A' <= first && first <= 'Z' // A-Z
}

// Check if string starts with lowercase letter
func StartsWithLower(text string) bool {
	first := text[0]
	return 'a' <= first && first <= 'z' // a-z
}

// Check if string starts with digit
func StartsWithDigit(text string) bool {
	first := text[0]
	return '0' <= first && first <= '9'
}

// Get leading whitespace
func SpacePrefix(text string) string {
	suffix := TrimLeftSpace(text)
	return strings.TrimSuffix(text, suffix)
}

// Get trailing whitespace
func SpaceSuffix(text string) string {
	prefix := TrimRightSpace(text)
	return strings.TrimPrefix(text, prefix)
}

// Trim left whitespace
func TrimLeftSpace(text string) string {
	return strings.TrimLeftFunc(text, unicode.IsSpace)
}

// Trim right whitespace
func TrimRightSpace(text string) string {
	return strings.TrimRightFunc(text, unicode.IsSpace)
}

// Try to get the nth part of the string split by separator
func TryGetPart(text, sep string, index int) (string, bool) {
	parts := strings.Split(text, sep)
	if len(parts) < index+1 {
		return "", false
	}
	return parts[index], true
}

// Wrap error with prefix error message
func WrapError(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}
