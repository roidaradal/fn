package check

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

func AllFalse(items []bool) bool {
	return AllEqual(items, false)
}

func Any[T any](items []T, ok func(T) bool) bool {
	for _, item := range items {
		if ok(item) {
			return true
		}
	}
	return false
}

func AnyEqual[T comparable](items []T, value T) bool {
	for _, item := range items {
		if item == value {
			return true
		}
	}
	return false
}

func AnyTrue(items []bool) bool {
	return AnyEqual(items, true)
}

func AnyFalse(items []bool) bool {
	return AnyEqual(items, false)
}
