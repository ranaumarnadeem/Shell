// ===== builtins/which.go =====
package builtins

import (
    "fmt"
    "io"
    "os/exec"
)


func Which(in io.Reader, out io.Writer, args []string, builtins []string) error {
    if len(args) == 0 {
        fmt.Fprintln(out, "Usage: which command")
        return nil
    }
    for _, cmd := range args {
        found := false
        for _, b := range builtins {
            if cmd == b {
                fmt.Fprintf(out, "%s: shell built-in\n", cmd)
                found = true
                break
            }
        }
        if !found {
            if path, err := exec.LookPath(cmd); err == nil {
                fmt.Fprintf(out, "%s: %s\n", cmd, path)
            } else {
                fmt.Fprintf(out, "%s: not found\n", cmd)
            }
        }
    }
    return nil
}

