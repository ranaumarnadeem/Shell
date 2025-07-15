package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SetEnvVar sets or lists environment variables
func SetEnvVar(in io.Reader, out io.Writer, args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(out, "Usage: setenv VAR [value]")
		return nil
	}
	varName := args[0]
	varValue := ""
	if len(args) > 1 {
		varValue = strings.Join(args[1:], " ")
	}
	os.Setenv(varName, varValue)
	fmt.Fprintf(out, "Set %s=%s\n", varName, varValue)
	return nil
}

// UnsetEnvVar removes an environment variable
func UnsetEnvVar(in io.Reader, out io.Writer, args []string) error {
	if len(args) < 1 {
		fmt.Fprintln(out, "Usage: unsetenv VAR")
		return nil
	}
	varName := args[0]
	os.Unsetenv(varName)
	fmt.Fprintf(out, "Unset %s\n", varName)
	return nil
}

func PrintEnv(in io.Reader, out io.Writer, args []string) error {
	for _, e := range os.Environ() {
		fmt.Fprintln(out, e)
	}
	return nil
}
