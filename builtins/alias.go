package builtins

import (
	"fmt"
	"strings"
)

func SetAlias(aliases map[string]string, args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: alias name command")
		return
	}
	name := args[0]
	command := strings.Join(args[1:], " ")
	aliases[name] = command
	fmt.Printf("alias set: %s='%s'\n", name, command)
}

func RemoveAlias(aliases map[string]string, args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: unalias name")
		return
	}
	delete(aliases, args[0])
	fmt.Printf("alias removed: %s\n", args[0])
}
