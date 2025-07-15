# 🥔 Potato Shell

A quirky, powerful custom shell written in Go — with support for built-in commands, aliasing, history, and a secret **Skibidi Mode** for real Rizzards! 🚽

---

## 🚀 Features

- 🧱 **Built-in Commands**: `ls`, `cd`, `echo`, `help`, `pwd`, etc.
- 🧠 **Command Aliasing**: `alias ll ls -la`
- 📜 **Command History**: Scroll through previously entered commands
- 🪄 **Skibidi Mode**: Type commands like `giga-walk`, `rizz-echo`, and `brainblast`
- 🪢 **Pipes Support**: Use `|` to chain commands
- 🌐 **External Commands**: Executes any binary available in `$PATH`
- 🧽 **Environment Interaction**: `env`, `setenv`, `unsetenv`
- 📂 **File Viewing**: `cat`, `clear`, `which`

---

## 🧑‍💻 Getting Started

```bash
go run main.go

## 🛠 Built-in Commands

| Command                  | Description                                     |
|--------------------------|-------------------------------------------------|
| `ls [-l] [-a] [dir]`     | List files (long/hidden supported)              |
| `cd <dir>`               | Change working directory                        |
| `pwd`                    | Print working directory                         |
| `cat <file>`             | Print contents of a file                        |
| `echo <args>`            | Print arguments to terminal                     |
| `clear`                  | Clear the terminal screen                       |
| `help`                   | Show help for built-in commands                 |
| `history`                | Show previously entered commands                |
| `alias <name> <cmd>`     | Define a new alias                              |
| `unalias <name>`         | Remove a defined alias                          |
| `which <cmd>`            | Show if command is built-in or system binary    |
| `env`                    | Show all environment variables                  |
| `setenv <k> <v>`         | Set an environment variable                     |
| `unsetenv <k>`           | Remove an environment variable                  |
| `skibidi-help`           | List Skibidi-mode command equivalents           |

---

## 🧌 Skibidi Mode

Potato Shell has a special **Skibidi Mode** for ultra-Rizzards.  
In this mode, common commands are remapped to ridiculous but fun aliases:

| Skibidi Command    | Normal Command |
|--------------------|----------------|
| `giga-walk`        | `cd`           |
| `skibidi-peek`     | `ls`           |
| `rizz-echo`        | `echo`         |
| `brainblast`       | `help`         |
| `old-tales`        | `history`      |
| `save-my-bits`     | `alias`        |
| `unskibidi`        | `unalias`      |
| `wheres-it-at`     | `which`        |
| `toxic-vars`       | `env`          |
| `spawn-var`        | `setenv`       |
| `nuke-var`         | `unsetenv`     |
| `cat-jam`          | `cat`          |
| `mirror-me`        | `pwd`          |
| `wipe-it`          | `clear`        |
| `skibidi-help`     | Shows this table 🧌 |

Use `skibidi-help` anytime in Skibidi Mode to see this again.

---

## 📝 Example Session

```bash
alias ll ls -la
ll
giga-walk ..
rizz-echo Yo, I’m in Skibidi Mode 🚽
save-my-bits greet echo Hello Rizz
greet


## 📦 Project Structure

```bash
main.go               # Main shell logic and REPL

/shell
  dispatcher.go       # Handles built-ins + Skibidi remapping
  executor.go         # Executes commands (internal/external)

/builtins
  ls.go, cd.go, ...   # Each built-in implemented separately
  skibidi_help.go     # Skibidi command help handler
## 📦 Dependencies

- ✅ **Standard Go library**
- [Optional] [`mvdan.cc/sh`](https://pkg.go.dev/mvdan.cc/sh) — For advanced POSIX-style shell parsing (not required by default)

---

## ⚙️ Requirements

- **Go**: 1.18 or higher

### OS Support

- ✅ **Linux** — Fully supported  
- 🪟 **Windows** — Partial support  
- 🍎 **macOS** — Basic support (not fully tested)

### External File Opening

- `xdg-open` (Linux)  
- `cmd /C start` (Windows)

---

## 🤓 Dev Notes

- Single-threaded by design  
- No job control (`&`, `fg`, `bg`) yet  
- Pipe support is basic (no redirection or subshells)  
- Skibidi Mode is toggleable **only at launch** (for now)

---

## ❤️ Inspired By

- Shells: **Bash**, **Zsh**, **Fish**
- Culture: **Memes**, **Rizz**, and too much time online
- Tech: Love for 🥔 **potatoes** and **Go**

---

## 📜 License

**MIT License**  
Do anything you want. Just **don't run**:

```bash
unalias rizz
