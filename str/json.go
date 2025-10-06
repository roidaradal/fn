package str

import (
	"encoding/json"
	"strings"
)

// Convert the item into a JSON string
func JSON[T any](item T) (string, error) {
	bytes, err := MarshalJSON(item, 0)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Convert the item into an indented JSON string
func IndentedJSON[T any](item T) (string, error) {
	bytes, err := MarshalJSON(item, 1)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Marshal the JSON item, with or without indent
func MarshalJSON[T any](item T, indent int) ([]byte, error) {
	if indent <= 0 {
		return json.Marshal(item)
	} else {
		return json.MarshalIndent(item, "", strings.Repeat("  ", indent))
	}
}
