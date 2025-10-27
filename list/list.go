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

// Applies the function to each item
func Decorate[T any](items []T, decorate func(T) T) []T {
	for i, item := range items {
		items[i] = decorate(item)
	}
	return items
}

func NumRange[T ~uint | ~int](start, end T) []T {
	items := make([]T, 0, end-start)
	for x := start; x < end; x++ {
		items = append(items, x)
	}
	return items
}
