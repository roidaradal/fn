package dict

import (
	"encoding/json"
	"maps"
	"slices"
)

type (
	StringMap     = map[string]string
	StringListMap = map[string][]string
	StringCounter = map[string]int
	IntCounter    = map[int]int
	Object        = map[string]any
)

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// Get map keys
func Keys[K comparable, V any](items map[K]V) []K {
	return slices.Collect(maps.Keys(items))
}

// Get map values
func Values[K comparable, V any](items map[K]V) []V {
	return slices.Collect(maps.Values(items))
}

// Get map entries
func Entries[K comparable, V any](items map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(items))
	for k, v := range items {
		entries = append(entries, Entry[K, V]{k, v})
	}
	return entries
}

// Check if map has key
func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, hasKey := items[key]
	return hasKey
}

// Set default value if key is not yet in map
func SetDefault[K comparable, V any](items map[K]V, key K, value V) {
	if _, ok := items[key]; !ok {
		items[key] = value
	}
}

// Zip the list of keys and values to form map
func Zip[K comparable, V any](keys []K, values []V) map[K]V {
	m := make(map[K]V, len(keys))
	numValues := len(values)
	for i, k := range keys {
		if i >= numValues {
			break // stop if no more values
		}
		m[k] = values[i]
	}
	return m
}

// Create a map from given struct pointer
func FromStruct[T any, V any](item *T) (map[string]V, error) {
	output := make(map[string]V)
	if item == nil {
		return output, nil
	}
	bytes, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
