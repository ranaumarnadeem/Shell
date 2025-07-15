package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"your-module-name/shell"
)

var (
	commandHistory []string
	aliases        = make(map[string]string)
	skibidiMode    bool
)

func main() {
	fmt.Println("🎉 Welcome to the Potato Shell!")

	// Mode selection
	fmt.Println("Choose your mode:")
	fmt.Println("1) Normal Mode (Boomer)")
	fmt.Println("2) Skibidi Mode (For Real Rizzards)")
	fmt.Print("Enter 1 or 2: ")

	var mode string
	fmt.Scanln(&mode)
	if mode == "2" {
		skibidiMode = true
		fmt.Println("🛸 Entering Skibidi Mode... GYATTTTT 🚽")
	} else {
		fmt.Println("👴 Normal Mode activated.")
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		shell.DisplayPrompt()

		if !scanner.Scan() {
			break
		}

		inputLine := strings.TrimSpace(scanner.Text())
		if inputLine == "" {
			continue
		}

		if inputLine == "exit" {
			if skibidiMode {
				fmt.Println("👋 Bye Bye Skibidi Bro.")
			} else {
				fmt.Println("Goodbye.")
			}
			break
		}

		commandHistory = append(commandHistory, inputLine)

		if strings.Contains(inputLine, "|") {
			err := shell.HandlePipes(inputLine, &commandHistory, &aliases, skibidiMode)
			if err != nil {
				fmt.Printf("Pipeline Error: %v\n", err)
			}
			continue
		}

		tokens, err := shell.ParseInput(inputLine, aliases)
		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}
		if tokens == nil || len(tokens) == 0 {
			continue
		}

		err = shell.HandleCommand(inputLine, aliases, commandHistory, skibidiMode)
		if err != nil {
			fmt.Printf("Command Error: %v\n", err)
		}
	}
}
