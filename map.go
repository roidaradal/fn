package fn

import (
	"maps"
	"slices"
)

func MapKeys[K comparable, V any](items map[K]V) []K {
	return slices.Collect(maps.Keys(items))
}

func MapValues[K comparable, V any](items map[K]V) []V {
	return slices.Collect(maps.Values(items))
}

func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, ok := items[key]
	return ok
}

func Zip[K comparable, V any](keys []K, values []V) map[K]V {
	m := make(map[K]V, len(keys))
	numValues := len(values)
	for i, k := range keys {
		if i >= numValues {
			break
		}
		m[k] = values[i]
	}
	return m
}
