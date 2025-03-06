package main

import (
	//"bufio"
	"fmt"
	"os"
)

func disp_path() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(dir)
	fmt.Print("> $")
}
func cd() {
	var path string
	//fmt.Print("Enter path: ")
	fmt.Scan(&path)
	err := os.Chdir(path)
	disp_path()
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	fmt.Println("Welcome to the shell")
	disp_path()
	//fmt.Print("> $")
	var command string
	fmt.Scan(&command)
	for command != "exit" {
		if command == "ls" {
			//ls()
			fmt.Println("will be added in future")
		}else if command == "cd" {
			cd()

		}
	//    disp_path()
	//	fmt.Print("> $")
		fmt.Scan(&command)
	}
}

