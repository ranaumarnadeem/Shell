package shell

import (
	"fmt"
	"io"
	"your-module-name/builtins"
)

var builtInList = []string{
	"cd", "ls", "echo", "help", "history", "alias", "unalias",
	"env", "setenv", "unsetenv", "which", "pwd", "clear", "cat", "skibidi-help",
}

func dispatchBuiltin(
	cmd string,
	in io.Reader,
	out io.Writer,
	args []string,
	aliases map[string]string,
	history []string,
	skibidiMode bool,
) error {
	if skibidiMode {
		cmd = SkibidiRemap(cmd, true)
	}
	fmt.Println("dispatching:", cmd, "skibidiMode:", skibidiMode)

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
	case "skibidi-help":
		return builtins.SkibidiHelp(out)

	default:
		return nil
	}
}

func SkibidiRemap(name string, skibidi bool) string {
	if !skibidi {
		return name
	}
	skibidiMap := map[string]string{
		"giga-walk":    "cd",
		"skibidi-peek": "ls",
		"rizz-echo":    "echo",
		"old-tales":    "history",
		"brainblast":   "help",
		"save-my-bits": "alias",
		"unskibidi":    "unalias",
		"wheres-it-at": "which",
		"toxic-vars":   "env",
		"nuke-var":     "unsetenv",
		"spawn-var":    "setenv",
		"cat-jam":      "cat",
		"mirror-me":    "pwd",
		"wipe-it":      "clear",
		"skibidi-help": "skibidi-help",
	}
	if normal, ok := skibidiMap[name]; ok {
		return normal
	}
	return name
}
