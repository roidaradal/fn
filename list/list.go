// Package list contains list-related functions.
package list

import (
	"math/rand/v2"
	"slices"
)

// Return list length
func Length[T any](items []T) int {
	return len(items)
}

// Returns new list with items copied from given list
func Copy[T any](items []T) []T {
	items2 := append([]T{}, items...)
	return items2
}

// Computes sum of items
func Sum[T ~uint | ~int | ~float32 | ~float64](items []T) T {
	var sum T = 0
	for _, item := range items {
		sum += item
	}
	return sum
}

// Computes product of items
func Product[T ~uint | ~int | ~float32 | ~float64](items []T) T {
	var product T = 1
	for _, item := range items {
		product *= item
	}
	return product
}

// Return list of integers from [start, end)
func NumRange[T ~uint | ~int](start, end T) []T {
	items := make([]T, 0, end-start)
	for x := start; x < end; x++ {
		items = append(items, x)
	}
	return items
}

// Shuffles the given items in-place
func Shuffle[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}

// Divides the number of items by the number of parts (divide workload).
// Return the index ranges of each part
func Divide(numItems, numParts int) [][2]int {
	portion := numItems / numParts
	ranges := make([][2]int, numParts)
	for i := range numParts - 1 {
		start := i * portion
		ranges[i] = [2]int{start, start + portion}
	}
	i := numParts - 1
	start := i * portion
	ranges[i] = [2]int{start, numItems}
	return ranges
}

// Tally the number of occurences of each item in the list
func TallyItems[T comparable](items []T) map[T]int {
	count := make(map[T]int)
	for _, item := range items {
		count[item] += 1
	}
	return count
}

// Count the number of occurences of given value in the list
func Count[T comparable](items []T, value T) int {
	count := 0
	for _, item := range items {
		if item == value {
			count += 1
		}
	}
	return count
}

// Creates a list containing the given <value> repeated <count> times
func Repeated[T any](value T, count int) []T {
	return slices.Repeat([]T{value}, count)
}
