package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"your-module-name/shell"
	"your-module-name/helper"
)

var commandHistory []string
var aliases = make(map[string]string)

func main() {
	fmt.Println("Welcome to the Potato Shell")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		shell.DisplayPrompt()

		if !scanner.Scan() {
			break
		}

		inputLine := strings.TrimSpace(scanner.Text())
		
		if inputLine == "exit" {
			fmt.Println("Bye Bye")
			break
		}
		
		if utils.IsEmpty(inputLine) {
			continue
		}
		
		commandHistory = append(commandHistory, inputLine)

		if strings.Contains(inputLine, "|") {
			err := shell.HandlePipes(inputLine, &commandHistory, &aliases)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
			continue
		}

		tokens, err := shell.ParseInput(inputLine, &aliases)
		if err != nil {
			fmt.Println(err)
			continue
		}
		
		if tokens == nil {
			continue
		}
		
		shell.ExecuteCommand(tokens, &commandHistory, &aliases)
	}
}