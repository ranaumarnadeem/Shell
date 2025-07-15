package builtins

import (
	"fmt"
	"io"
	"os"
)

// Cd changes the current working directory
func Cd(in io.Reader, out io.Writer, args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(out, "cd: missing operand")
		return nil
	}
	path := args[0]
	if err := os.Chdir(path); err != nil {
		fmt.Fprintf(out, "cd error: %v\n", err)
	}
	return nil
}
