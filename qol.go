package fn

func Ternary[T any](condition bool, valueTrue T, valueFalse T) T {
	if condition {
		return valueTrue
	} else {
		return valueFalse
	}
}
