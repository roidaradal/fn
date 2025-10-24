package str

import (
	"fmt"
	"strings"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/dyn"
)

// Convert *string to string
func RefString(item *string) string {
	if item == nil {
		return ""
	}
	return strings.TrimSpace(*item)
}

// Convert int to string
func Int[T ~int | ~uint](x T) string {
	return fmt.Sprintf("%d", x)
}

// Convert *int to string
func RefInt[T ~int | ~uint](x *T) string {
	if x == nil {
		return ""
	}
	return Int(*x)
}

// Convert float to string
func Float[T ~float32 | ~float64](x T) string {
	return fmt.Sprintf("%f", x)
}

// Convert Boolean to string (0, 1)
func Boolean(flag bool) string {
	return fn.Ternary(flag, "1", "0")
}

// Convert string to Boolean (false if "0", true otherwise)
func ToBoolean(flag string) bool {
	return flag != "0"
}

// Convert string to *string, nil if empty string
func ToRefString(item string) *string {
	item = strings.TrimSpace(item)
	return fn.Ternary(item == "", nil, &item)
}

// Ensure *string contains non-empty string, nil otherwise
func NonEmptyRefString(item *string) *string {
	if item == nil {
		return nil
	}
	return ToRefString(*item)
}

// Convert any to string
func Any(item any) string {
	if dyn.IsPointer(item) {
		return fmt.Sprintf("%v", dyn.Deref(item))
	}
	return fmt.Sprintf("%v", item)
}
