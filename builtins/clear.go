// ===== builtins/clear.go =====
package builtins

import (
	"fmt"
	"io"
)

func Clear(in io.Reader, out io.Writer, args []string) error {

	fmt.Fprint(out, "\033[H\033[2J")
	return nil
}
