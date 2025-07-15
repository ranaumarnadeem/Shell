package builtins

import (
	"fmt"
	"io"
)

// Help displays available built-in commands and usage
func Help(in io.Reader, out io.Writer, args []string) error {
	fmt.Fprintln(out, "Potato Shell - Built-in Commands:")
	fmt.Fprintln(out, "______________________________________________________________________________")
	fmt.Fprintln(out, "|  ls [-l] [-a] [dir]  |   List files")
	fmt.Fprintln(out, "|  cd <dir>            |   Change directory")
	fmt.Fprintln(out, "|  open <file>         |   Open a file")
	fmt.Fprintln(out, "|  pwd                 |   Print working directory")
	fmt.Fprintln(out, "|  clear               |   Clear the screen")
	fmt.Fprintln(out, "|  help                |   Show this help message")
	fmt.Fprintln(out, "|  exit                |   Exit the shell")
	fmt.Fprintln(out, "|  echo <text>         |   Print text")
	fmt.Fprintln(out, "|  history             |   Show command history")
	fmt.Fprintln(out, "|  alias [name cmd]    |   Create or list aliases")
	fmt.Fprintln(out, "|  unalias <name>      |   Remove an alias")
	fmt.Fprintln(out, "|  which <cmd>         |   Locate command")
	fmt.Fprintln(out, "|  setenv VAR [value]  |   Set environment variable")
	fmt.Fprintln(out, "|  unsetenv VAR        |   Unset environment variable")
	fmt.Fprintln(out, "|  env                 |   List environment variables")
	fmt.Fprintln(out, "______________________________________________________________________________")
	return nil
}
