package builtins

import "fmt"

func ShowHistory(history []string) {
	for i, cmd := range history {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}
