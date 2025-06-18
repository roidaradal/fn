package fn

import (
	"maps"
	"slices"
)

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

func MapKeys[K comparable, V any](items map[K]V) []K {
	return slices.Collect(maps.Keys(items))
}

func MapValues[K comparable, V any](items map[K]V) []V {
	return slices.Collect(maps.Values(items))
}

func MapEntries[K comparable, V any](items map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(items))
	for k, v := range items {
		entries = append(entries, Entry[K, V]{k, v})
	}
	return entries
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
