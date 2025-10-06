package check

// Check if all list items pass the ok function
func All[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if !ok(item) {
			return false
		}
	}
	return true
}

// Check if any list item passes the ok function
func Any[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

// Check if all list items are equal to the given value
func AllEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item != value {
			return false
		}
	}
	return true
}

// Check if any list item is equal to the given value
func AnyEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

// Check if all list items are true
func AllTrue(items []bool) bool {
	return AllEqual(items, true)
}

// Check if all list items are false
func AllFalse(items []bool) bool {
	return AllEqual(items, false)
}

// Check if any list item is true
func AnyTrue(items []bool) bool {
	return AnyEqual(items, true)
}

// Check if any list item is false
func AnyFalse(items []bool) bool {
	return AnyEqual(items, false)
}
