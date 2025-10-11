package dyn

// Check if two any values are equal
func Equal(item1 any, item2 any) bool {
	// Dereference item1 if pointer and not null
	if IsPointer(item1) && !IsNull(item1) {
		item1 = Deref(item1)
	}
	// Dereference item2 if pointer and not null
	if IsPointer(item2) && !IsNull(item2) {
		item2 = Deref(item2)
	}
	return item1 == item2
}

// Check if two any values are not equal
func NotEqual(item1 any, item2 any) bool {
	return !Equal(item1, item2)
}
