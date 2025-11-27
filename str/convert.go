package str

import (
	"fmt"
	"strings"
)

// Convert item to string
func Any[T any](item T) string {
	return fmt.Sprintf("%v", item)
}

// Convert list of items to string list
func List[T any](items []T) []string {
	items2 := make([]string, len(items))
	for i, item := range items {
		items2[i] = Any(item)
	}
	return items2
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

// Convert *float to string
func RefFloat[T ~float32 | ~float64](x *T) string {
	if x == nil {
		return ""
	}
	return Float(*x)
}

// Convert Boolean to string (0, 1)
func Boolean(flag bool) string {
	if flag {
		return "1"
	}
	return "0"
}

// Convert string to Boolean (false if "0", true otherwise)
func ToBoolean(flag string) bool {
	return flag != "0"
}

// Convert *string to string
func RefString(item *string) string {
	if item == nil {
		return ""
	}
	return strings.TrimSpace(*item)
}

// Convert string to *string, nil if empty string
func ToRefString(item string) *string {
	item = strings.TrimSpace(item)
	if item == "" {
		return nil
	}
	return &item
}

// Ensure *string contains non-empty string, nil otherwise
func NonEmptyRefString(item *string) *string {
	if item == nil {
		return nil
	}
	return ToRefString(*item)
}

// Convert to '.' if empty string
func GuardDot(item string) string {
	if item == "" {
		return "."
	}
	return item
}

// Convert to '.' if empty string, and convert to uppercase
func UpperDot(item string) string {
	return strings.ToUpper(GuardDot(item))
}

// Convert to '.' if empty string, and convert to lowercase
func LowerDot(item string) string {
	return strings.ToLower(GuardDot(item))
}
