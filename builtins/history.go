package builtins

import (
	"fmt"
	"io"
)

func ShowHistory(in io.Reader, out io.Writer, history []string) error {
	for i, cmd := range history {
		fmt.Fprintf(out, "%d: %s\n", i+1, cmd)
	}
	return nil
}
