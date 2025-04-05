package fn

func RunForever() {
	for {
		select {}
	}
}

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
