// Package dict contains map-related functions.
package dict

import (
	"cmp"
	"maps"
	"slices"
)

// Get map length
func Length[K comparable, V any](items map[K]V) int {
	return len(items)
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

// Get map entries, sorted by keys
func SortedEntries[K cmp.Ordered, V any](items map[K]V) []Entry[K, V] {
	keys := Keys(items)
	slices.Sort(keys)
	entries := make([]Entry[K, V], len(keys))
	for i, k := range keys {
		entries[i] = Entry[K, V]{k, items[k]}
	}
	return entries
}

// Sort the list of values for each key
func SortValues[K comparable, V cmp.Ordered](items map[K][]V) {
	for k, values := range items {
		slices.Sort(values)
		items[k] = values
	}
}

// Check if map has key
func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, hasKey := items[key]
	return hasKey
}

// Check if map has value
func HasValue[K, V comparable](items map[K]V, value V) bool {
	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}

// Check if map has no key
func NoKey[K comparable, V any](items map[K]V, key K) bool {
	return !HasKey(items, key)
}

// Check if map has no value
func NoValue[K, V comparable](items map[K]V, value V) bool {
	return !HasValue(items, value)
}

// Set default value if key is not in map
func SetDefault[K comparable, V any](items map[K]V, key K, defaultValue V) {
	if _, ok := items[key]; !ok {
		items[key] = defaultValue
	}
}

// Get the value of key, or return default value if key is not in map
func DefaultGet[K comparable, V any](items map[K]V, key K, defaultValue V) V {
	if value, ok := items[key]; ok {
		return value
	}
	return defaultValue
}

// Get value = obj[key], then type coerce into T
func Get[T any](obj Object, key string) (T, bool) {
	var item T
	value, ok := obj[key]
	if !ok {
		return item, false
	}
	item, ok = value.(T)
	return item, ok
}

// Get value = obj[key], then type coerce into *T
func GetRef[T any](obj Object, key string) *T {
	itemRef, ok := Get[*T](obj, key)
	if !ok {
		return nil
	}
	return itemRef
}

// Get value = obj[key] then type coerce into []*T
func GetListRef[T any](obj Object, key string) []*T {
	listRef, ok := Get[[]*T](obj, key)
	if !ok {
		return nil
	}
	return listRef
}
