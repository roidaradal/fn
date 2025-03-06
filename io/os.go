package io

import (
	"os"
	"path/filepath"
)

const defaultFileMode = 0o666

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func EnsurePathExists(path string) error {
	return os.MkdirAll(filepath.Dir(path), defaultFileMode)
}
