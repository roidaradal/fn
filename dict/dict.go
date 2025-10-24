package dict

import (
	"cmp"
	"encoding/json"
	"maps"
	"slices"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/dyn"
)

type (
	UintMap       = map[string]uint
	IntMap        = map[string]int
	BoolMap       = map[string]bool
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

// Get map entries, sorted by keys
func SortedEntries[K cmp.Ordered, V any](items map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(items))
	keys := Keys(items)
	slices.Sort(keys)
	for _, k := range keys {
		entries = append(entries, Entry[K, V]{k, items[k]})
	}
	return entries
}

// Check if map has key
func HasKey[K comparable, V any](items map[K]V, key K) bool {
	_, hasKey := items[key]
	return hasKey
}

// Check if map has value
func HasValue[K comparable, V comparable](items map[K]V, value V) bool {
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
func NoValue[K comparable, V comparable](items map[K]V, value V) bool {
	return !HasValue(items, value)
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

// Create an Object from given struct pointer
func ToObject[T any](item *T) (Object, error) {
	return FromStruct[T, any](item)
}

// Create a struct from given object
func ToStruct[T any](item Object) (*T, error) {
	var output T
	if item == nil {
		return &output, nil
	}
	bytes, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Add the entries of new map into old map, return old map
func Update[K comparable, V any](oldMap map[K]V, newMap map[K]V) map[K]V {
	for k, v := range newMap {
		oldMap[k] = v
	}
	return oldMap
}

// Convert a struct to an Object, but only keeping the given fieldNames
func Prune[T any](structRef *T, fieldNames ...string) *Object {
	object := make(Object)
	for _, fieldName := range fieldNames {
		value := dyn.GetFieldValue(structRef, fieldName)
		object[fieldName] = value
	}
	return &object
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
	return fn.Ternary(ok, itemRef, nil)
}

// Get value = obj[key], then type coerce into []*T
func GetListRef[T any](obj Object, key string) []*T {
	listRef, ok := Get[[]*T](obj, key)
	return fn.Ternary(ok, listRef, nil)
}
