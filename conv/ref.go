package conv

import (
	"strings"

	"github.com/roidaradal/fn"
)

func Ref[T any](item T) *T {
	return &item
}

func Deref[T any](item *T) T {
	return *item
}

func NullIfBlank(item string) *string {
	item = strings.TrimSpace(item)
	return fn.Ternary(item == "", nil, &item)
}

func Nullable(item *string) *string {
	if item == nil {
		return nil
	}
	return NullIfBlank(*item)
}

func NullToString(item *string) string {
	if item == nil {
		return ""
	}
	return strings.TrimSpace(*item)
}

func NullToBlank(item *string) string {
	item2 := NullToString(item)
	return fn.Ternary(item2 == "", "-", item2)
}

func NullIntToString[T ~int | ~uint](item *T) string {
	if item == nil {
		return ""
	}
	return IntToString(*item)
}
