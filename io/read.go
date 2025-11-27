package io

import (
	"encoding/json"
	"os"

	"github.com/roidaradal/fn/list"
	"github.com/roidaradal/fn/str"
)

// Read string contents of given text file path
func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Read lines of given text file path
func ReadLines(path string) ([]string, error) {
	text, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	return str.Lines(text), nil
}

// Read non-empty lines of given text file path
func ReadNonEmptyLines(path string) ([]string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	lines = list.Filter(lines, str.NotEmpty)
	return lines, nil
}

// Read rows of given CSV file path
func ReadCSV(path string) ([][]string, error) {
	lines, err := ReadNonEmptyLines(path)
	if err != nil {
		return nil, err
	}
	rows := list.Map(lines, str.CommaSplit)
	return rows, nil
}

// Read JSON object from given file path
func ReadJSON[T any](path string) (*T, error) {
	item, err := readJSON[T](path)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Read JSON list from given file path
func ReadJSONList[T any](path string) ([]T, error) {
	return readJSON[[]T](path)
}

// Read JSON map from given file path
func ReadJSONMap[T any](path string) (map[string]T, error) {
	return readJSON[map[string]T](path)
}

// Common: unmarshal JSON from file
func readJSON[T any](path string) (T, error) {
	var item T
	bytes, err := os.ReadFile(path)
	if err != nil {
		return item, err
	}
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return item, err
	}
	return item, nil
}
