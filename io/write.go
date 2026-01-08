package io

import (
	"os"

	"github.com/roidaradal/fn/str"
)

// Save string to given file path
func SaveString(text, path string) error {
	return os.WriteFile(path, []byte(text), defaultFileMode)
}

// Append string to given file path
func AppendString(text, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, defaultFileMode)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text)
	return err
}

// Save JSON object to given file path
func SaveJSON[T any](item T, path string) error {
	return saveJSON(item, path, 0)
}

// Save indented JSON object to given file path
func SaveIndentedJSON[T any](item T, path string) error {
	return saveJSON(item, path, 1)
}

// Common: save JSON to file path, with or without indent
func saveJSON[T any](item T, path string, indent int) error {
	bytes, err := str.MarshalJSON(item, indent)
	if err != nil {
		return err
	}
	err = EnsurePathExists(path)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bytes, defaultFileMode)
}
