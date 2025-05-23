package helper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen clears the terminal/console screen.
// It works cross-platform by using the appropriate command based on the operating system:
// - Windows: uses "cls" command
// - Unix/Linux/macOS: uses "clear" command
// If the command execution fails, it falls back to using ANSI escape sequences.
func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// For Linux, macOS, etc.
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()

	// Fallback to ANSI escape sequence if command execution fails
	if err != nil {
		fmt.Print("\033[H\033[2J")
	}
}
