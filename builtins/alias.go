package builtins

import (
	"fmt"
	"io"
	"strings"
)

// SetAlias creates a new alias or lists existing ones
func SetAlias(in io.Reader, out io.Writer, aliases map[string]string, args []string) error {
	if len(args) == 0 {
		// list all aliases
		for name, cmd := range aliases {
			fmt.Fprintf(out, "alias %s='%s'\n", name, cmd)
		}
		return nil
	}
	if len(args) < 2 {
		fmt.Fprintln(out, "Usage: alias name command")
		return nil
	}
	name := args[0]
	command := strings.Join(args[1:], " ")
	if strings.Contains(name, " ") {
		fmt.Fprintln(out, "Invalid alias name")
		return nil
	}
	aliases[name] = command
	fmt.Fprintf(out, "alias set: %s='%s'\n", name, command)
	return nil
}

// RemoveAlias deletes an existing alias
func RemoveAlias(in io.Reader, out io.Writer, aliases map[string]string, args []string) error {
	if len(args) < 1 {
		fmt.Fprintln(out, "Usage: unalias name")
		return nil
	}
	delete(aliases, args[0])
	fmt.Fprintf(out, "alias removed: %s\n", args[0])
	return nil
}
