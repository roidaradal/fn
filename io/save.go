package io

import (
	"encoding/json"
	"os"
	"strings"
)

func StringifyJSON[T any](item T) (string, error) {
	data, err := stringifyJSON(item, 0)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func StringifyIndentedJSON[T any](item T) (string, error) {
	data, err := stringifyJSON(item, 1)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SaveJSONList[T any](path string, items []T) error {
	return saveJSON(path, items, 0)
}

func SaveIndentedJSONList[T any](path string, items []T) error {
	return saveJSON(path, items, 1)
}

func SaveJSONMap[T any](path string, items map[string]T) error {
	return saveJSON(path, items, 0)
}

func SaveIndentedJSONMap[T any](path string, items map[string]T) error {
	return saveJSON(path, items, 1)
}

func SaveJSONObject[T any](path string, obj T) error {
	return saveJSON(path, obj, 0)
}

func SaveIndentedJSONObject[T any](path string, obj T) error {
	return saveJSON(path, obj, 1)
}

func stringifyJSON[T any](item T, indent int) ([]byte, error) {
	if indent == 0 {
		return json.Marshal(item)
	} else {
		return json.MarshalIndent(item, "", strings.Repeat("  ", indent))
	}
}

func saveJSON[T any](path string, item T, indent int) error {
	data, err := stringifyJSON(item, indent)
	if err != nil {
		return err
	}
	err = EnsurePathExists(path)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
