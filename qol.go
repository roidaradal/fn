package fn

func Ternary[T any](condition bool, valueTrue T, valueFalse T) T {
	if condition {
		return valueTrue
	} else {
		return valueFalse
	}
}

func Deref[T any](item *T) T {
	return *item
}

func Ref[T any](item T) *T {
	return &item
}
