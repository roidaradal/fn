// Package io contains input/output and filesystem-related functions.
package io

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

// Opens file using default viewer
func OpenFile(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}
	return cmd.Start()
}
