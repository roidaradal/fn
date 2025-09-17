package io

import (
	"encoding/json"
	"os"

	"github.com/roidaradal/fn"
	"github.com/roidaradal/fn/check"
	"github.com/roidaradal/fn/str"
)

func ReadTextFile(path string) (string, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(text), nil
}

func ReadAllTextLines(path string) ([]string, error) {
	text, err := ReadTextFile(path)
	if err != nil {
		return nil, err
	}
	lines := str.CleanSplit(text, "\n")
	return lines, nil
}

func ReadTextLines(path string) ([]string, error) {
	lines, err := ReadAllTextLines(path)
	if err != nil {
		return nil, err
	}
	lines = fn.Filter(lines, check.IsNotBlankString)
	return lines, nil
}

func ReadCSVFile(path string) ([][]string, error) {
	lines, err := ReadTextLines(path)
	if err != nil {
		return nil, err
	}
	rows := make([][]string, len(lines))
	for i, line := range lines {
		rows[i] = str.CleanSplit(line, ",")
	}
	return rows, nil
}

func LoadJSONObject[T any](path string) (*T, error) {
	obj, err := loadJSON[T](path)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

func LoadJSONList[T any](path string) ([]T, error) {
	return loadJSON[[]T](path)
}

func LoadJSONMap[T any](path string) (map[string]T, error) {
	return loadJSON[map[string]T](path)
}

func loadJSON[T any](path string) (T, error) {
	var obj T
	data, err := os.ReadFile(path)
	if err != nil {
		return obj, err
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return obj, err
	}
	return obj, nil
}
