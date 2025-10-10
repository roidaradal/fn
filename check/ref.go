package check

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
