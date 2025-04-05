package fn

import (
	"strings"
)

func IsBlankString(text string) bool {
	return text == ""
}

func IsNotBlankString(text string) bool {
	return text != ""
}

func CleanSplit(text string, sep string) []string {
	parts := strings.Split(text, sep)
	parts = Map(parts, strings.TrimSpace)
	return parts
}

func SpaceSplit(text string) []string {
	parts := strings.Fields(strings.TrimSpace(text))
	parts = Map(parts, strings.TrimSpace)
	return parts
}

func NullIfBlank(item string) *string {
	item = strings.TrimSpace(item)
	if item == "" {
		return nil
	} else {
		return &item
	}
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
	if item2 == "" {
		return "-"
	}
	return item2
}
