package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"your-module-name/builtins"

	"mvdan.cc/sh/v3/shell"
)

var commandHistory []string
var aliases = make(map[string]string)

// List of built-in commands for the which command
var builtinCommands = []string{"cd", "ls", "help", "alias", "unalias", "history", "echo", "open", "exit", "which", "setenv", "unsetenv", "env"}

func displayPrompt() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s> $ ", dir)
}

func main() {
	fmt.Println("Welcome to the Potato Shell")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		displayPrompt()

		if !scanner.Scan() {
			break
		}

		inputLine := strings.TrimSpace(scanner.Text())
		if inputLine == "exit" {
			fmt.Println("Bye Bye")
			break
		}

		if inputLine == "" {
			continue
		}

		commandHistory = append(commandHistory, inputLine)

		// Expand environment variables
		expandedInput := os.ExpandEnv(inputLine)

		tokens, err := shell.Fields(expandedInput, nil)
		if err != nil {
			fmt.Println("Error parsing command:", err)
			continue
		}

		if len(tokens) == 0 {
			continue
		}

		// Alias expansion
		if aliasCmd, ok := aliases[tokens[0]]; ok {
			newInput := aliasCmd + " " + strings.Join(tokens[1:], " ")
			tokens = strings.Fields(newInput)
		}

		command := tokens[0]
		args := tokens[1:]

		switch command {
		case "ls":
			builtins.Ls(args)
		case "cd":
			builtins.Cd(args)
		case "open":
			builtins.OpenFile(args)
		case "help":
			builtins.Help()
		case "history":
			builtins.ShowHistory(commandHistory)
		case "alias":
			builtins.SetAlias(aliases, args)
		case "unalias":
			builtins.RemoveAlias(aliases, args)
		case "echo":
			builtins.Echo(args)
		case "which":
			builtins.Which(args, builtinCommands)
		case "setenv":
			builtins.SetEnvVar(args)
		case "unsetenv":
			builtins.UnsetEnvVar(args)
		case "env":
			builtins.PrintEnv()
		default:
			cmdPath, err := exec.LookPath(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
				continue
			}

			cmd := exec.Command(cmdPath, args...)
			cmd.Env = os.Environ()
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err = cmd.Run()
			if err != nil {
				fmt.Printf("Error running command: %v\n", err)
			}
		}
	}
}
