package builtins

import (
	"fmt"
	"os/exec"
)

func Which(args []string, builtins []string) {
	if len(args) == 0 {
		fmt.Println("which: missing operand")
		return
	}

	for _, cmd := range args {
		// Check if it's a built-in command
		isBuiltin := false
		for _, b := range builtins {
			if b == cmd {
				fmt.Printf("%s: shell builtin\n", cmd)
				isBuiltin = true
				break
			}
		}
		if isBuiltin {
			continue
		}

		// Check in PATH
		path, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Printf("%s: not found\n", cmd)
		} else {
			fmt.Println(path)
		}
	}
}
