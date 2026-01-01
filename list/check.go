package list

import (
	"cmp"
	"slices"
)

// Check if list is empty
func IsEmpty[T any](items []T) bool {
	return len(items) == 0
}

// Check if list is not empty
func NotEmpty[T any](items []T) bool {
	return len(items) > 0
}

// Check if all list items pass the ok function
func All[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if !ok(item) {
			return false
		}
	}
	return true
}

// Check if any list item passes the ok function
func Any[T any](items []T, ok func(T) bool) bool {
	return slices.ContainsFunc(items, ok)
}

// Check if all list items pass the ok function
// (accepts index and item)
func IndexedAll[T any](items []T, ok func(int, T) bool) bool {
	for i, item := range items {
		if !ok(i, item) {
			return false
		}
	}
	return true
}

// Check if any list item passes the ok function
// (accepts index and item)
func IndexedAny[T any](items []T, ok func(int, T) bool) bool {
	for i, item := range items {
		if ok(i, item) {
			return true
		}
	}
	return false
}

// Check if all list items are equal to the given value
func AllEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item != value {
			return false
		}
	}
	return true
}

// Check if any list item is equal to the given value
func AnyEqual[T comparable](items []T, value T) bool {
	return slices.Contains(items, value)
}

// Check if all list items are true
func AllTrue(items []bool) bool {
	return AllEqual(items, true)
}

// Check if all list items are false
func AllFalse(items []bool) bool {
	return AllEqual(items, false)
}

// Check if any list item is true
func AnyTrue(items []bool) bool {
	return AnyEqual(items, true)
}

// Check if any list item is false
func AnyFalse(items []bool) bool {
	return AnyEqual(items, false)
}

// Check if all list items are the same
func AllSame[T comparable](items []T) bool {
	return len(TallyItems(items)) == 1
}

// Check if all list items are unique
func AllUnique[T comparable](items []T) bool {
	return len(TallyItems(items)) == len(items)
}

// Check if all list items are not equal to the given value
func AllNotEqual[T comparable](items []T, value T) bool {
	return !slices.Contains(items, value)
}

// Check if all list items are greater than given value
func AllGreater[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x > value
	})
}

// Check if all list items are greater or equal than given value
func AllGreaterEqual[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x >= value
	})
}

// Check if all list items are lesser than given value
func AllLess[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x < value
	})
}

// Check if all list items are lesser or equal than given value
func AllLessEqual[T cmp.Ordered](items []T, value T) bool {
	return All(items, func(x T) bool {
		return x <= value
	})
}
