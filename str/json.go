package str

import (
	"encoding/json"
	"strings"
)

// Create JSON string from struct
func JSON[T any](item T) (string, error) {
	bytes, err := MarshalJSON(item, 0)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Create indented JSON string from struct
func IndentedJSON[T any](item T) (string, error) {
	bytes, err := MarshalJSON(item, 1)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Marshal struct as JSON, with or without indent
func MarshalJSON[T any](item T, indent int) ([]byte, error) {
	if indent <= 0 {
		return json.Marshal(item)
	} else {
		return json.MarshalIndent(item, "", strings.Repeat("  ", indent))
	}
}
