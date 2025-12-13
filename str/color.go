package str

import "fmt"

const reset string = "\033[0m"

// Wraps text in red color
func Red(text string) string {
	return fmt.Sprintf("\033[31m%s%s", text, reset)
}

// Wraps text in green color
func Green(text string) string {
	return fmt.Sprintf("\033[32m%s%s", text, reset)
}

// Wraps text in yellow color
func Yellow(text string) string {
	return fmt.Sprintf("\033[33m%s%s", text, reset)
}

// Wraps text in blue color
func Blue(text string) string {
	return fmt.Sprintf("\033[34m%s%s", text, reset)
}

// Wraps text in violet color
func Violet(text string) string {
	return fmt.Sprintf("\033[35m%s%s", text, reset)
}

// Wraps text in cyan color
func Cyan(text string) string {
	return fmt.Sprintf("\033[36m%s%s", text, reset)
}
