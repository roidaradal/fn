package io

import (
	"os"
	"path/filepath"
)

const defaultFileMode os.FileMode = 0o666

// Check if path is a directory
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}

// Check if path exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Creates all non-existent folders in the given path
func EnsurePathExists(path string) error {
	return os.MkdirAll(filepath.Dir(path), defaultFileMode)
}
