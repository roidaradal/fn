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

// Unzip the map, return the list of keys and values,
// Order of keys is same as the order of corresponding values
func Unzip[K comparable, V any](items map[K]V) ([]K, []V) {
	numItems := len(items)
	keys := make([]K, 0, numItems)
	values := make([]V, 0, numItems)
	for k, v := range items {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
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

// Add the entries of new map into old map, return old map
func Update[K comparable, V any](oldMap map[K]V, newMap map[K]V) map[K]V {
	for k, v := range newMap {
		oldMap[k] = v
	}
	return oldMap
}
