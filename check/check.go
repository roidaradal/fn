// Package check contains validation and checker functions.
package check

import (
	"net/mail"
	"regexp"
)

// Checks if string is alphanumeric
func IsAlphanumeric(text string) bool {
	alphanumeric := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return alphanumeric.MatchString(text)
}

// Validate email address
func IsValidEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}

// Check if pointer is nil
func IsNull[T any](item *T) bool {
	return item == nil
}

// Check if pointer is not nil
func NotNull[T any](item *T) bool {
	return item != nil
}

// Compares two references if their underlying values are equal
func RefValueEqual[T comparable](item1 *T, item2 *T) bool {
	if item1 == nil && item2 == nil {
		// both nil = equal
		return true
	} else if item1 != nil && item2 != nil {
		// both not nil = compare deref values
		return *item1 == *item2
	} else {
		// one is nil, other is not = not equal
		return false
	}
}

// Compares two references if their underlying values are not equal
func RefValueNotEqual[T comparable](item1 *T, item2 *T) bool {
	return !RefValueEqual(item1, item2)
}

// Creates function that checks if value is equal to goal value
func IsEqual[T comparable](goalValue T) func(T) bool {
	return func(value T) bool {
		return value == goalValue
	}
}

// Creates function that checks if value is not equal to goal value
func NotEqual[T comparable](goalValue T) func(T) bool {
	return func(value T) bool {
		return value != goalValue
	}
}
