package io

import (
	"os"

	"github.com/roidaradal/fn/str"
)

// Save JSON object to given file path
func SaveJSON[T any](item T, path string) error {
	return saveJSON(item, path, 0)
}

// Save indented JSON object to given file path
func SaveIndentedJSON[T any](item T, path string) error {
	return saveJSON(item, path, 1)
}

// Internal: save JSON to file, with or without indent
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
