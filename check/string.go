package check

import "net/mail"

func IsBlankString(text string) bool {
	return text == ""
}

func IsNotBlankString(text string) bool {
	return text != ""
}

func IsValidEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}
