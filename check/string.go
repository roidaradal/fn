package check

import (
	"net/mail"
	"regexp"
)

var alphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")

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

// Checks if string is alphanumeric
func IsAlphaNumeric(text string) bool {
	return alphaNumeric.MatchString(text)
}
