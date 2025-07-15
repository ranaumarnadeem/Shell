// builtins/cat.go
package builtins

import (
	"io"
)

// Cat copies everything from stdin to stdout.
func Cat(in io.Reader, out io.Writer, args []string) error {
	_, err := io.Copy(out, in)
	return err
}
