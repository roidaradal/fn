package list

// Apply convert function to each list item, return mapped items
func Map[T, V any](items []T, convert func(T) V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = convert(item)
	}
	return results
}

// Apply convert function to each list item with index, return mapped items
func IndexedMap[T, V any](items []T, convert func(int, T) V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = convert(i, item)
	}
	return results
}

// Map list indexes to given values; return mapped items.
// Can have zero value if index is invalid
func MapList[T any](indexes []int, values []T) []T {
	numValues := len(values)
	results := make([]T, len(indexes))
	for i, idx := range indexes {
		if 0 <= idx && idx < numValues {
			results[i] = values[idx]
		}
	}
	return results
}

// Apply the translation mask to each list item, return mapped items
func Translate[K comparable, V any](items []K, mask map[K]V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = mask[item]
	}
	return results
}

// Filter list: only keep items that pass keep function
func Filter[T any](items []T, keep func(T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			results = append(results, item)
		}
	}
	return results
}

// Filter list: only keep items with index that pass keep function
func IndexedFilter[T any](items []T, keep func(int, T) bool) []T {
	results := make([]T, 0, len(items))
	for i, item := range items {
		if keep(i, item) {
			results = append(results, item)
		}
	}
	return results
}

// Apply task function to each item
func Apply[T any](items []T, task func(T) T) []T {
	for i, item := range items {
		items[i] = task(item)
	}
	return items
}

// Deduplicate list, preserving given order
func Deduplicate[T comparable](items []T) []T {
	done := make(map[T]bool)
	unique := make([]T, 0, len(items))
	for _, item := range items {
		if done[item] {
			continue
		}
		unique = append(unique, item)
		done[item] = true
	}
	return unique
}
