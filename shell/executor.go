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
func HandleCommand(cmdStr string, aliases map[string]string, history []string, skibidiMode bool) error {
	tokens, err := ParseInput(cmdStr, aliases)
	if err != nil {
		return err
	}
	if len(tokens) == 0 {
		return nil
	}

	name := tokens[0]
	args := tokens[1:]

	if isBuiltin(name, skibidiMode) {
		return dispatchBuiltin(name, os.Stdin, os.Stdout, args, aliases, history, skibidiMode)
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

// AddHistory appends a command to the global history
func AddHistory(cmd string) {
	history = append(history, cmd)
}

// Updated isBuiltin to support Skibidi remapping
func isBuiltin(cmd string, skibidiMode bool) bool {
	if skibidiMode {
		cmd = SkibidiRemap(cmd, true)
	}
	for _, b := range builtInList {
		if cmd == b {
			return true
		}
	}
	return false
}
