package builtins

import "fmt"

func Help() {
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
	fmt.Println("|  setenv VAR [value]  |   Set environment variable")
	fmt.Println("|  unsetenv VAR        |   Remove environment variable")
	fmt.Println("|  env                 |   Show all environment variables")
	fmt.Println("______________________________________________________________________________")
}
