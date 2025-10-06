package fn

// Shortcut for ternary operator: condition ? valueTrue : valueFalse
func Ternary[T any](condition bool, valueTrue T, valueFalse T) T {
	if condition {
		return valueTrue
	} else {
		return valueFalse
	}
}

// Apply the convert function to each list item, return the resulting items
func Map[T any, S any](items []T, convert func(T) S) []S {
	results := make([]S, len(items))
	for i, item := range items {
		results[i] = convert(item)
	}
	return results
}

// Apply the convert function to each list item with index, return the resulting items
func MapIndex[T any, S any](items []T, convert func(int, T) S) []S {
	results := make([]S, len(items))
	for i, item := range items {
		results[i] = convert(i, item)
	}
	return results
}

// Apply the translation map to each list item, return the resulting items
func Translate[K comparable, V any](items []K, mask map[K]V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = mask[item]
	}
	return results
}

// Filter out list items that fail the keep function
func Filter[T any](items []T, keep func(T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			results = append(results, item)
		}
	}
	return results
}

// Filter out list items with index that fail the keep function
func FilterIndex[T any](items []T, keep func(int, T) bool) []T {
	results := make([]T, 0, len(items))
	for i, item := range items {
		if keep(i, item) {
			results = append(results, item)
		}
	}
	return results
}

// Filter out map entries that fail that keep function
func FilterMap[K comparable, V any](items map[K]V, keep func(K, V) bool) map[K]V {
	results := make(map[K]V)
	for k, v := range items {
		if keep(k, v) {
			results[k] = v
		}
	}
	return results
}
