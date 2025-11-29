package lang

import (
	"testing"
)

func TestTernary(t *testing.T) {
	type testCase[T any] struct {
		condition  bool
		valueTrue  T
		valueFalse T
		expected   T
	}
	testCases := []testCase[int]{
		{true, 1, 0, 1},
		{false, 1, 0, 0},
	}
	for _, x := range testCases {
		output := Ternary(x.condition, x.valueTrue, x.valueFalse)
		if output != x.expected {
			t.Errorf("exp: %v, got: %v", x.expected, output)
		}
	}
}

func TestDeref(t *testing.T) {
	type testCase[T any] struct {
		ref      *T
		expected T
	}
	a, b := 1, 2
	testCases := []testCase[int]{
		{&a, a},
		{&b, b},
		{nil, 0},
	}
	for _, x := range testCases {
		output := Deref(x.ref)
		if output != x.expected {
			t.Errorf("exp: %v, got: %v", x.expected, output)
		}
	}
}
