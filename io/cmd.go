package io

import (
	"fmt"
	"os"
	"os/exec"
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
