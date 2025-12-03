package dict

import (
	"encoding/json"
	"maps"
)

// Add entries of new map into old map, returns old map.
// If there are key conflicts, new map entries overwrite the old map entries.
func Update[K comparable, V any](oldMap, newMap map[K]V) map[K]V {
	maps.Copy(oldMap, newMap)
	return oldMap
}

// Zip list of keys and values to create map
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

// Unzip the map, return list of keys and values,
// Order of keys is same as order of corresponding values
func Unzip[K comparable, V any](items map[K]V) ([]K, []V) {
	numItems := len(items)
	keys := make([]K, numItems)
	values := make([]V, numItems)
	i := 0
	for k, v := range items {
		keys[i] = k
		values[i] = v
		i++
	}
	return keys, values
}

// Swap keys and values, convert map[K]V to map[V]K,
// Can lose data if values are not unique
func Swap[K, V comparable](items map[K]V) map[V]K {
	inverse := make(map[V]K, len(items))
	for k, v := range items {
		inverse[v] = k
	}
	return inverse
}

// Convert map[K][]V to map[V]K,
// Can lose data if values are not unique
func SwapList[K, V comparable](items map[K][]V) map[V]K {
	inverse := make(map[V]K)
	for k, values := range items {
		for _, v := range values {
			inverse[v] = k
		}
	}
	return inverse
}

// Group data by values
func GroupByValue[K, V comparable](items map[K]V) map[V][]K {
	groups := make(map[V][]K)
	for k, v := range items {
		groups[v] = append(groups[v], k)
	}
	return groups
}

// Group data (map[K][]V) by values => map[V][]K
func GroupByValueList[K, V comparable](items map[K][]V) map[V][]K {
	groups := make(map[V][]K)
	for k := range items {
		for _, v := range items[k] {
			groups[v] = append(groups[v], k)
		}
	}
	return groups
}

// Create Object from given struct pointer
func ToObject[T any](structRef *T) (Object, error) {
	return FromStruct[T, any](structRef)
}

// Create Object from struct, but only keep given fieldNames
func Prune[T any](structRef *T, fieldNames ...string) (Object, error) {
	fullObj, err := ToObject(structRef)
	if err != nil {
		return nil, err
	}
	obj := make(Object)
	for _, fieldName := range fieldNames {
		if value, ok := fullObj[fieldName]; ok {
			obj[fieldName] = value
		}
	}
	return obj, nil
}

// Create map from given struct pointer
func FromStruct[T, V any](structRef *T) (map[string]V, error) {
	output := make(map[string]V)
	if structRef == nil {
		return output, nil
	}
	bytes, err := json.Marshal(structRef)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Create struct from given Object
func ToStruct[T any](obj Object) (*T, error) {
	var output T
	if obj == nil {
		return &output, nil
	}
	bytes, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &output)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Create tally of how many times each value appears in the map
func TallyValues[K, V comparable](items map[K]V, values []V) map[V]int {
	count := NewCounter(values)
	for _, value := range items {
		if NoKey(count, value) {
			continue // skip value if not in counter
		}
		count[value] += 1
	}
	return count
}

// Filter map: only keep entries that pass keep function
func Filter[K comparable, V any](items map[K]V, keep func(K, V) bool) map[K]V {
	results := make(map[K]V, len(items))
	for k, v := range items {
		if keep(k, v) {
			results[k] = v
		}
	}
	return results
}

// Merge the counts from counter maps into one map
func MergeCounts[K comparable](counters []map[K]int) map[K]int {
	total := make(map[K]int)
	for _, counter := range counters {
		for key, count := range counter {
			total[key] += count
		}
	}
	return total
}
