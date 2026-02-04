// Package io contains input/output and filesystem-related functions.
package io

import (
	"os"
	"path/filepath"
	"strings"
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

// Creates all non-existent folders in given path
func EnsurePathExists(path string) error {
	return os.MkdirAll(filepath.Dir(path), defaultFileMode)
}

// Gets the filename without extension
func BaseFileName(path string) string {
	filename := filepath.Base(path)
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
