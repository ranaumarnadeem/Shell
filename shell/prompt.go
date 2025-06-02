package shell

import (
	"fmt"
	"os"
)

// DisplayPrompt shows the current directory prompt
func DisplayPrompt() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%s> $ ", dir)
}
