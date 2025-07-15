package shell

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Stage represents one element of the pipeline,
// with Run taking an input reader and an output writer.
type Stage struct {
	Run func(in io.Reader, out io.Writer) error
}

func HandlePipes(input string) error {
	parts := strings.Split(input, "|")
	if len(parts) < 2 {
		return fmt.Errorf("invalid pipe syntax: %s", input)
	}

	var stages []Stage
	for _, part := range parts {
		cmdStr := strings.TrimSpace(part)
		if cmdStr == "" {
			return fmt.Errorf("empty command in pipeline: %s", input)
		}

		tokens, err := ParseInput(cmdStr)
		if err != nil {
			return fmt.Errorf("parse error for '%s': %w", cmdStr, err)
		}
		if len(tokens) == 0 {
			return fmt.Errorf("no tokens in segment: %s", cmdStr)
		}

		name, args := tokens[0], tokens[1:]
		if isBuiltin(name) {
			// Built-in stage
			stages = append(stages, Stage{
				Run: func(in io.Reader, out io.Writer) error {
					return dispatchBuiltin(name, in, out, args, aliases, builtInList)
				},
			})
		} else {
			// External command stage
			stages = append(stages, NewExternalStage(name, args))
		}
	}

	return ExecuteStages(stages)
}

func ExecuteStages(stages []Stage) error {
	n := len(stages)
	readers := make([]io.Reader, n)
	writers := make([]io.WriteCloser, n)

	for i := 0; i < n-1; i++ {
		r, w := io.Pipe()
		writers[i] = w
		readers[i+1] = r
	}

	readers[0] = os.Stdin
	writers[n-1] = nopWriteCloser{os.Stdout}

	for i, st := range stages {
		in := readers[i]
		out := writers[i]
		go func(s Stage, in io.Reader, out io.Writer) {
			defer func() {
				if wc, ok := out.(io.WriteCloser); ok {
					wc.Close()
				}
			}()

			if err := s.Run(in, out); err != nil {
				fmt.Fprintf(os.Stderr, "pipeline stage error: %v\n", err)
			}
		}(st, in, out)
	}

	return nil
}

// nopWriteCloser wraps an io.Writer to satisfy io.WriteCloser (no-op Close)
type nopWriteCloser struct{ io.Writer }

// Close does nothing for the wrapper
func (nopWriteCloser) Close() error { return nil }
