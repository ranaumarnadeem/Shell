package builtins

import (
	"fmt"
	"os/exec"
	"runtime"
)

func OpenFile(args []string) {
	if len(args) == 0 {
		fmt.Println("open: missing file operand")
		return
	}

	file := args[0]
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", file)
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", "", file)
	case "darwin":
		cmd = exec.Command("open", file)
	default:
		fmt.Println("Unsupported OS")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
}
