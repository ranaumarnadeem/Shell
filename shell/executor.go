package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// history stores the list of commands entered by the user
var history []string

func NewExternalStage(name string, args []string) Stage {
	return Stage{
		Run: func(in io.Reader, out io.Writer) error {
			cmd := exec.Command(name, args...)
			cmd.Stdin = in
			cmd.Stdout = out
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	}
}

// HandleCommand executes a single (non-piped) command string
// It parses input, checks built-ins, or invokes an external command
func HandleCommand(input string) error {
	tokens, err := ParseInput(input)
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		return nil
	}

	name := tokens[0]
	args := tokens[1:]

	if isBuiltin(name) {
		// Run built-in with access to stdin/stdout
		// Pass empty environment variables and redirects to match the expected signature
		return dispatchBuiltin(name, os.Stdin, os.Stdout, args, map[string]string{}, []string{})
	}

	
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
	
		if ee, ok := err.(*exec.Error); ok && ee.Err == exec.ErrNotFound {
			return fmt.Errorf("%s: command not found", name)
		}
		return fmt.Errorf("%s: %v", name, err)
	}
	return nil
}


func AddHistory(cmd string) {
	history = append(history, cmd)
}
