// Package lang contains quality-of-life functions.
package lang

// Shortcut for ternary operator: condition ? valueTrue : valueFalse
func Ternary[T any](condition bool, valueTrue T, valueFalse T) T {
	if condition {
		return valueTrue
	} else {
		return valueFalse
	}
}

// Returns reference to given item
func Ref[T any](item T) *T {
	return &item
}

// Dereference given item pointer,
// If null pointer, return zero value of item
func Deref[T any](ref *T) T {
	var item T
	if ref != nil {
		item = *ref
	}
	return item
}
