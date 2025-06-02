package builtins

import (
	"fmt"
	"os"
)

func Cd(args []string) {
	if len(args) == 0 {
		fmt.Println("cd: missing operand")
		return
	}
	for _, path := range args {
		err := os.Chdir(path)
		if err != nil {
			fmt.Println("cd error:", err)
		}
	}
}
