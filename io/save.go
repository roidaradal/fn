package io

import (
	"encoding/json"
	"os"
	"strings"
)

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

func saveJSON[T any](path string, items T, indent int) error {
	var data []byte
	var err error
	if indent == 0 {
		data, err = json.Marshal(items)
	} else {
		data, err = json.MarshalIndent(items, "", strings.Repeat("  ", indent))
	}
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
