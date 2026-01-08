package io

import "fmt"

// Clear screen
func ClearScreen() {
	fmt.Print("\033[2J\033[H")
	// cmd := exec.Command("cmd", "/c", "cls")
	// cmd.Stdout = os.Stdout
	// cmd.Run()
}
