package io

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Clear screen
func ClearScreen() {
	command := "clear"
	if runtime.GOOS == "windows" {
		command = "cls"
	}
	RunCommand(command)
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

// Run command line command
func RunCommand(args ...string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		args = append([]string{"/c"}, args...)
		cmd = exec.Command("cmd", args...)
	default:
		args = append([]string{"-c"}, args...)
		cmd = exec.Command("sh", args...)
	}
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Run go install <path>
func RunGoInstall(path string) error {
	cmd1, cmd2 := "go", "install"
	fmt.Printf("Running: %s %s %s ... ", cmd1, cmd2, path)
	err := RunCommand(cmd1, cmd2, path)
	if err != nil {
		fmt.Println("FAIL")
		return err
	}
	fmt.Println("OK")
	return nil
}

// Get command and options
func GetCommandOptions(defaultCommand string) (string, map[string]string) {
	args := os.Args[1:]
	if len(args) < 1 {
		args = append(args, defaultCommand)
	}
	command := strings.ToLower(args[0])
	options := make(map[string]string)
	for _, pair := range args[1:] {
		parts := strings.SplitN(pair, "=", 2)
		key := strings.ToLower(parts[0])
		value := ""
		if len(parts) == 2 {
			value = parts[1]
		}
		options[key] = value
	}
	return command, options
}
