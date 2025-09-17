package fn

func Map[T any, S any](items []T, convert func(T) S) []S {
	results := make([]S, len(items))
	for i, item := range items {
		results[i] = convert(item)
	}
	return results
}

func MapIndex[T any, S any](items []T, convert func(int, T) S) []S {
	results := make([]S, len(items))
	for i, item := range items {
		results[i] = convert(i, item)
	}
	return results
}

func Filter[T any](items []T, keep func(T) bool) []T {
	results := make([]T, 0, len(items))
	for _, item := range items {
		if keep(item) {
			results = append(results, item)
		}
	}
	return results
}

func FilterIndex[T any](items []T, keep func(int, T) bool) []T {
	results := make([]T, 0, len(items))
	for i, item := range items {
		if keep(i, item) {
			results = append(results, item)
		}
	}
	return results
}

func FilterMap[K comparable, V any](items map[K]V, keep func(K, V) bool) map[K]V {
	items2 := make(map[K]V)
	for k, v := range items {
		if keep(k, v) {
			items2[k] = v
		}
	}
	return items2
}

func Ternary[T any](condition bool, valueTrue T, valueFalse T) T {
	if condition {
		return valueTrue
	} else {
		return valueFalse
	}
}

func RunForever() {
	for {
		select {}
	}
}
