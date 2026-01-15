package io

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Clear screen
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Soft Clear Screen
func SoftClearScreen() {
	fmt.Print("\033[2J\033[H")
}

// Opens file using default viewer
func OpenFile(path string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// Add "" for quote problems
		cmd = exec.Command("cmd", "/c", "start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	default:
		cmd = exec.Command("xdg-open", path)
	}
	return cmd.Start()
}
