// ===== builtins/pwd.go =====
package builtins

import (
    "fmt"
    "io"
    "os"
)


func Pwd(in io.Reader, out io.Writer, args []string) error {
    dir, err := os.Getwd()
    if err != nil {
        fmt.Fprintf(out, "pwd error: %v\n", err)
        return nil
    }
    fmt.Fprintln(out, dir)
    return nil
}
