package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func HandlePipes(input string, history *[]string, aliases *map[string]string) error {
	commands := strings.Split(input, "|")
	if len(commands) < 2 {
		return fmt.Errorf("too short input : %s", input)
	}

	var cmds []*exec.Cmd
	for _, cmdStr := range commands {
		cmdStr = strings.TrimSpace(cmdStr)
		if cmdStr == "" {
			return fmt.Errorf("empty command")
		}

		tokens, err := ParseInput(cmdStr, aliases)
		if err != nil {
			return fmt.Errorf("can't parse the pipe command: %w", err)
		}

		if len(tokens) == 0 {
			return fmt.Errorf("empty token")
		}

		cmd := exec.Command(tokens[0], tokens[1:]...)
		cmds = append(cmds, cmd)
	}

	return executePipe(cmds)
}

func executePipe(cmds []*exec.Cmd) error {

	for i := 0; i < len(cmds)-1; i++ {
		r, w := io.Pipe()
		cmds[i].Stdout = w
		cmds[i+1].Stdin = r
	}

	lastCmd := cmds[len(cmds)-1]
	lastCmd.Stdout = os.Stdout
	lastCmd.Stderr = os.Stderr

	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("pipe start error: %w", err)
		}
	}

	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return fmt.Errorf("pipe wait error: %w", err)
		}
	}

	return nil
}
