package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

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

func main() {
	fmt.Println("Welcome to the shell")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		dispPath()

		if !scanner.Scan() {
			break // Exit the loop if there's an error or EOF
		}
		inputLine := strings.TrimSpace(scanner.Text())
		if inputLine == "exit" {
			break
		}

		tokens := strings.Fields(inputLine)
		if len(tokens) == 0 {
			continue
		}
		command := tokens[0]
		args := tokens[1:]
		switch command {
		case "ls":
			ls(args)
		case "cd":
			cd(args)
		default:
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
