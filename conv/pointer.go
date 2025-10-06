package conv

// Returns a reference to given item
func Ref[T any](item T) *T {
	return &item
}

// De-references the given item pointer,
// If null pointer, returns the zero value of item
func Deref[T any](ref *T) T {
	var item T
	if ref != nil {
		item = *ref
	}
	return item
}
