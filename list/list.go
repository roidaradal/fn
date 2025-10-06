package list

// Returns a new list with the items copied from the given list
func Copy[T any](items []T) []T {
	items2 := append([]T{}, items...)
	return items2
}

// Computes the sum of items
func Sum[T ~uint | ~int | ~float32 | ~float64](items []T) T {
	var sum T = 0
	for _, item := range items {
		sum += item
	}
	return sum
}

// Computes the product of items
func Product[T ~uint | ~int | ~float32 | ~float64](items []T) T {
	var product T = 1
	for _, item := range items {
		product *= item
	}
	return product
}
