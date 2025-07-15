package builtins

import (
	"fmt"
	"io"
)

func SkibidiHelp(out io.Writer) error {
	fmt.Fprintln(out, "🚽 Skibidi Mode Command Mappings 🚽")
	fmt.Fprintln(out, "----------------------------------")
	mapper := map[string]string{
		"giga-walk":    "cd",
		"skibidi-peek": "ls",
		"rizz-echo":    "echo",
		"old-tales":    "history",
		"brainblast":   "help",
		"save-my-bits": "alias",
		"unskibidi":    "unalias",
		"wheres-it-at": "which",
		"toxic-vars":   "env",
		"spawn-var":    "setenv",
		"nuke-var":     "unsetenv",
	}
	for skibidi, normal := range mapper {
		fmt.Fprintf(out, "%-15s => %s\n", skibidi, normal)
	}
	return nil
}
