package dict

// Swap: map[K]V to map[V]K, can lose data if values are not unique
func Swap[K comparable, V comparable](items map[K]V) map[V]K {
	inverse := make(map[V]K)
	for k, v := range items {
		inverse[v] = k
	}
	return inverse
}

// Swap: map[K][]V to map[V]K, can lose data if values are not unique
func SwapList[K comparable, V comparable](items map[K][]V) map[V]K {
	inverse := make(map[V]K)
	for k, values := range items {
		for _, v := range values {
			inverse[v] = k
		}
	}
	return inverse
}

// Group data by values
func GroupByValue[K comparable, V comparable](items map[K]V) map[V][]K {
	groups := make(map[V][]K)
	for k, v := range items {
		groups[v] = append(groups[v], k)
	}
	return groups
}
