// ===== builtins/ls.go =====
package builtins

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type fileEntry struct {
	Name string
	Info os.FileInfo
}

func humanReadable(size int64) string {
	units := []string{"B", "K", "M", "G", "T"}
	idx := 0
	sz := float64(size)
	for sz >= 1024 && idx < len(units)-1 {
		sz /= 1024
		idx++
	}
	return fmt.Sprintf("%.1f%s", sz, units[idx])
}

func Ls(in io.Reader, out io.Writer, args []string) error {
	fs := flag.NewFlagSet("ls", flag.ContinueOnError)
	fs.SetOutput(out)

	showAll := fs.Bool("a", false, "Include hidden files")
	longList := fs.Bool("l", false, "Long listing format")
	human := fs.Bool("h", false, "Human-readable file sizes")
	reverse := fs.Bool("r", false, "Reverse order")
	sortTime := fs.Bool("t", false, "Sort by modification time (desc)")

	if err := fs.Parse(args); err != nil {
		return err
	}

	dir := "."
	if fs.NArg() > 0 {
		dir = fs.Arg(0)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	var files []fileEntry
	for _, entry := range entries {
		if !*showAll && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		files = append(files, fileEntry{entry.Name(), info})
	}

	if *sortTime {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Info.ModTime().After(files[j].Info.ModTime())
		})
	} else {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name < files[j].Name
		})
	}
	if *reverse {
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
	}

	for _, f := range files {
		if *longList {
			size := fmt.Sprintf("%d", f.Info.Size())
			if *human {
				size = humanReadable(f.Info.Size())
			}
			fmt.Fprintf(out, "%-10s %6s %s %s\n",
				f.Info.Mode().String(),
				size,
				f.Info.ModTime().Format("Jan 2 15:04"),
				f.Name)
		} else {
			fmt.Fprintln(out, f.Name)
		}
	}

	return nil
}
