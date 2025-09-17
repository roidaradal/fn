package list

func Copy[T any](items []T) []T {
	items2 := append([]T{}, items...)
	return items2
}

func Sum[T ~int | ~float32 | ~float64](items []T) T {
	var sum T = 0
	for _, item := range items {
		sum += item
	}
	return sum
}

func Product[T ~int | ~float32 | ~float64](items []T) T {
	var product T = 1
	for _, item := range items {
		product *= item
	}
	return product
}

func Lookup[T any, K comparable, V any](items []T, entry func(T) (K, V)) map[K]V {
	lookup := make(map[K]V, len(items))
	for _, item := range items {
		k, v := entry(item)
		lookup[k] = v
	}
	return lookup
}

func Translate[K comparable, V any](items []K, mask map[K]V) []V {
	results := make([]V, len(items))
	for i, item := range items {
		results[i] = mask[item]
	}
	return results
}
