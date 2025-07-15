package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"your-module-name/shell"
)

func main() {
	fmt.Println("Welcome to the Potato Shell")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		shell.DisplayPrompt()

		if !scanner.Scan() {
			fmt.Println()
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if input == "exit" {
			fmt.Println("Bye Bye")
			os.Exit(0)
		}

		shell.AddHistory(input)

		if strings.Contains(input, "|") {
			if err := shell.HandlePipes(input); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
			continue
		}

		if err := shell.HandleCommand(input); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}
}
