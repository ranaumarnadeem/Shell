package shell

import (
	"io"
	"your-module-name/builtins"
)

var builtInList = []string{
	"cd", "ls", "echo", "help", "history", "alias", "unalias",
	"env", "setenv", "unsetenv", "which", "pwd", "clear", "cat",
}

func isBuiltin(cmd string) bool {
	for _, b := range builtInList {
		if cmd == b {
			return true
		}
	}
	return false
}

func dispatchBuiltin(cmd string, in io.Reader, out io.Writer, args []string, aliases map[string]string, history []string) error {
	switch cmd {
	case "cd":
		return builtins.Cd(in, out, args)
	case "ls":
		return builtins.Ls(in, out, args)
	case "echo":
		return builtins.Echo(in, out, args)
	case "help":
		return builtins.Help(in, out, args)
	case "history":
		return builtins.ShowHistory(in, out, history)
	case "alias":
		return builtins.SetAlias(in, out, aliases, args)
	case "unalias":
		return builtins.RemoveAlias(in, out, aliases, args)
	case "env":
		return builtins.PrintEnv(in, out, args)
	case "setenv":
		return builtins.SetEnvVar(in, out, args)
	case "unsetenv":
		return builtins.UnsetEnvVar(in, out, args)
	case "which":
		return builtins.Which(in, out, args, builtInList)
	case "pwd":
		return builtins.Pwd(in, out, args)
	case "clear":
		return builtins.Clear(in, out, args)
	case "cat":
        return builtins.Cat(in, out, args)	
	default:
		return nil
	}
}

func ShowHistory(in io.Reader, out io.Writer, history []string) error {
	// Implementation of ShowHistory
	return nil
}
