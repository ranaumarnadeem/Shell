// builtins/echo.go
package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(in io.Reader, out io.Writer, args []string) error {

	interpretEscapes := false
	if len(args) > 0 && args[0] == "-e" {
		interpretEscapes = true
		args = args[1:]
	}

	s := strings.Join(args, " ")
	if interpretEscapes {
		s = strings.ReplaceAll(s, `\n`, "\n")
		s = strings.ReplaceAll(s, `\t`, "\t")
	}

	fmt.Fprintln(out, s)
	return nil
}
