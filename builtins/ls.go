package builtins

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Ls(args []string) {
	lsFlags := flag.NewFlagSet("ls", flag.ContinueOnError)
	showLong := lsFlags.Bool("l", false, "Use long listing format")
	showAll := lsFlags.Bool("a", false, "Show hidden files")

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
