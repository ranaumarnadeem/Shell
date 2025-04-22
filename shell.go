package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"mvdan.cc/sh/v3/shell"
)

var commandHistory []string
var aliases = make(map[string]string)

func help() {
	fmt.Println("Potato Shell - Built-in Commands:")
	fmt.Println("______________________________________________________________________________")
	fmt.Println("|  ls [-l] [-a] [dir]  |   List files in the current or specified directory")
	fmt.Println("|      -l              |   Use a long listing format")
	fmt.Println("|      -a              |   Show hidden files")
	fmt.Println("|  cd <dir>            |   Change the current directory")
	fmt.Println("|  open <file>         |   Open a file with the default associated program")
	fmt.Println("| help                 |   Show this help message")
	fmt.Println("|  exit                |   Exit the shell")
	fmt.Println("|  echo				|	Does whatever echo do")
	fmt.Println("|  history             |   Show past commands")
	fmt.Println("|  alias name command  |   Create a shortcut")
	fmt.Println("|  unalias name        |   Remove a shortcut")
	fmt.Println("|  which				|	Search for file or command in system path")
	fmt.Println("______________________________________________________________________________")
}

func showHistory() {
	for i, cmd := range commandHistory {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}

func setAlias(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: alias name command")
		return
	}
	name := args[0]
	command := strings.Join(args[1:], " ")
	aliases[name] = command
	fmt.Printf("alias set: %s='%s'\n", name, command)
}

func removeAlias(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: unalias name")
		return
	}
	delete(aliases, args[0])
	fmt.Printf("alias removed: %s\n", args[0])
}

func openFile(args []string) {
	if len(args) == 0 {
		fmt.Println("open: missing file operand")
		return
	}

	file := args[0]
	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "linux":
		cmd = exec.Command("xdg-open", file)
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", "", file)
	default:
		fmt.Println("Unsupported OS")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
}

func dispPath() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s> $ ", dir)
}

func cd(args []string) {
	if len(args) == 0 {
		fmt.Println("cd: missing operand")
		return
	}
	for _, path := range args {
		err := os.Chdir(path)
		if err != nil {
			fmt.Println("cd error:", err)
		}
	}
}

func ls(args []string) {
	lsFlags := flag.NewFlagSet("ls", flag.ContinueOnError)
	showLong := lsFlags.Bool("l", false, "garam")
	showAll := lsFlags.Bool("a", false, "anday.")

	err := lsFlags.Parse(args)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	remainingArgs := lsFlags.Args()
	dir := "."
	if len(remainingArgs) > 0 {
		dir = remainingArgs[0]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("ls error:", err)
		return
	}

	for _, file := range files {
		if !*showAll && strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if *showLong {
			info, err := file.Info()
			if err != nil {
				fmt.Println(file.Name())
			} else {
				fmt.Printf("%v %6d %v %s\n", info.Mode(), info.Size(), info.ModTime().Format("Jan 2 15:04"), file.Name())
			}
		} else {
			fmt.Println(file.Name())
		}
	}
}
func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}
func isBuiltIn(cmd string) bool {
	builtIns := []string{"cd", "ls", "help", "alias", "unalias", "history", "echo", "open", "exit", "which"}
	for _, b := range builtIns {
		if b == cmd {
			return true
		}
	}
	return false
}

func which(args []string) {
	if len(args) == 0 {
		fmt.Println("which: missing operand")
		return
	}

	for _, cmd := range args {
		// If it's a built-in shell command, just say so
		if isBuiltIn(cmd) {
			fmt.Printf("%s: shell builtin\n", cmd)
			continue
		}

		// Check in PATH for executable
		path, err := exec.LookPath(cmd)
		if err != nil {
			fmt.Printf("%s: not found\n", cmd)
		} else {
			fmt.Println(path)
		}
	}
}

func main() {
	fmt.Println("Welcome to the Potato Shell")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		dispPath()

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

		tokens, err := shell.Fields(inputLine, nil)
if err != nil {
	fmt.Println("Error parsing command:", err)
	continue
}
		if len(tokens) == 0 {
			continue
		}

		if aliasCmd, ok := aliases[tokens[0]]; ok {

			inputLine = aliasCmd + " " + strings.Join(tokens[1:], " ")
			tokens = strings.Fields(inputLine)
		}

		command := tokens[0]
		args := tokens[1:]

		switch command {
		case "ls":
			ls(args)
		case "cd":
			cd(args)
		case "open":
			openFile(args)
		case "help":
			help()
		case "history":
			showHistory()
		case "alias":
			setAlias(args)
		case "unalias":
			removeAlias(args)
		case "echo":
			echo(args)
		case "which":
			which(args)
		default:
			cmdPath, err := exec.LookPath(command)
			if err != nil {
				fmt.Printf("%s: command not found\n", command)
				continue
			}
		
			
			cmd := exec.Command(cmdPath, args...)
		
			
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
