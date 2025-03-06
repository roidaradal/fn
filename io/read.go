package io

import (
	"encoding/json"
	"os"

	"github.com/roidaradal/fn"
)

func ReadTextFile(path string) (string, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(text), nil
}

func ReadTextLines(path string) ([]string, error) {
	text, err := ReadTextFile(path)
	if err != nil {
		return nil, err
	}
	lines := fn.CleanSplit(text, "\n")
	lines = fn.Filter(lines, fn.IsNotBlankString)
	return lines, nil
}

func ReadCSVFile(path string) ([][]string, error) {
	lines, err := ReadTextLines(path)
	if err != nil {
		return nil, err
	}
	rows := make([][]string, len(lines))
	for i, line := range lines {
		rows[i] = fn.CleanSplit(line, ",")
	}
	return rows, nil
}

func LoadJSONObject[T any](path string) (*T, error) {
	var obj T
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
