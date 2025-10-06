package check

import "net/mail"

// Check empty string, used for filter functions
func IsEmptyString(text string) bool {
	return text == ""
}

// Check non-empty string, used for filter functions
func NotEmptyString(text string) bool {
	return text != ""
}

// Validate email address
func IsValidEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}
