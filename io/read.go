package io

import (
	"encoding/json"
	"os"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/check"
	"github.com/roidaradal/fn/str"
)

// Read contents of the given text file path
func ReadFile(path string) (string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Read all lines of given text file path
func ReadAllLines(path string) ([]string, error) {
	text, err := ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := str.Lines(text)
	return lines, nil
}

// Read non-blank lines of given text file path
func ReadLines(path string) ([]string, error) {
	lines, err := ReadAllLines(path)
	if err != nil {
		return nil, err
	}
	lines = fn.Filter(lines, check.NotEmptyString)
	return lines, nil
}

// Read rows of given CSV file path
func ReadCSV(path string) ([][]string, error) {
	lines, err := ReadLines(path)
	if err != nil {
		return nil, err
	}
	rows := make([][]string, len(lines))
	for i, line := range lines {
		rows[i] = str.CleanSplit(line, ",")
	}
	return rows, nil
}

// Read JSON object from given file path
func ReadJSON[T any](path string) (*T, error) {
	obj, err := readJSON[T](path)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// Read JSON list from given file path
func ReadJSONList[T any](path string) ([]T, error) {
	return readJSON[[]T](path)
}

// Read JSON map from given file path
func ReadJSONMap[T any](path string) (map[string]T, error) {
	return readJSON[map[string]T](path)
}

// Internal: unmarshal JSON from file
func readJSON[T any](path string) (T, error) {
	var obj T
	bytes, err := os.ReadFile(path)
	if err != nil {
		return obj, err
	}
	err = json.Unmarshal(bytes, &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
