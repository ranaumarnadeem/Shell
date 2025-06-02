package builtins

import (
	"fmt"
	"os"
	"strings"
)

func SetEnvVar(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: setenv VAR [value]")
		return
	}

	varName := args[0]
	varValue := ""
	if len(args) > 1 {
		varValue = strings.Join(args[1:], " ")
	}

	os.Setenv(varName, varValue)
	fmt.Printf("Set %s=%s\n", varName, varValue)
}

func UnsetEnvVar(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: unsetenv VAR")
		return
	}

	varName := args[0]
	os.Unsetenv(varName)
	fmt.Printf("Unset %s\n", varName)
}

func PrintEnv() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
