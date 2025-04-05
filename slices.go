package fn

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func CopySlice[T any](items []T) []T {
	items2 := append([]T{}, items...)
	return items2
}

func All[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if !ok(item) {
			return false
		}
	}
	return true
}

func AllEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item != value {
			return false
		}
	}
	return true
}

func AllTrue(items []bool) bool {
	return AllEqual(items, true)
}

func Any[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

func Sum[T Number](items []T) T {
	var sum T = 0
	for _, item := range items {
		sum += item
	}
	return sum
}

func Product[T Number](items []T) T {
	var product T = 1
	for _, item := range items {
		product *= item
	}
	return product
}
