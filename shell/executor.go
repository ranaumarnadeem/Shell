package shell

import (
	"fmt"
	"os"
	"os/exec"

	"your-module-name/builtins"
)

// this handles both built-in and external command execution
func ExecuteCommand(tokens []string, history *[]string, aliases *map[string]string) {
	if len(tokens) == 0 {
		return
	}

	command := tokens[0]
	args := tokens[1:]

	switch command {
	case "cd":
		builtins.Cd(args)
	case "ls":
		builtins.Ls(args)
	case "open":
		builtins.OpenFile(args)
	case "help":
		builtins.Help()
	case "history":
		builtins.ShowHistory(*history)
	case "alias":
		builtins.SetAlias(*aliases, args)
	case "unalias":
		builtins.RemoveAlias(*aliases, args)
	case "echo":
		builtins.Echo(args)
	case "which":
		builtinCommands := []string{"cd", "ls", "help", "alias", "unalias", "history", "echo", "open", "exit", "which", "setenv", "unsetenv", "env"}
		builtins.Which(args, builtinCommands)
	case "setenv":
		builtins.SetEnvVar(args)
	case "unsetenv":
		builtins.UnsetEnvVar(args)
	case "env":
		builtins.PrintEnv()
	case "exit":
		fmt.Println("Bye Bye")
		os.Exit(0)
	default:
		executeExternal(command, args)
	}
}

// executeExternal runs external commands
func executeExternal(command string, args []string) {
	cmdPath, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("%s: command not found\n", command)
		return
	}

	cmd := exec.Command(cmdPath, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
