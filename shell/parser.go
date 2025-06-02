package shell

import (
	"fmt"
	"os"
	"strings"

	"mvdan.cc/sh/v3/shell"
)

//processes the input line into tokens
func ParseInput(input string, aliases *map[string]string) ([]string, error) {
	
	expandedInput := os.ExpandEnv(input)

	// Parse into tokens
	tokens, err := shell.Fields(expandedInput, nil)
	if err != nil {
		return nil, fmt.Errorf("error parsing command: %w", err)
	}

	if len(tokens) == 0 {
		return nil, nil
	}

	// Handle alias expansion
	return expandAliases(tokens, aliases), nil
}

// replaces commands with their aliased equivalents
func expandAliases(tokens []string, aliases *map[string]string) []string {
	if aliasCmd, ok := (*aliases)[tokens[0]]; ok {
		newInput := aliasCmd
		if len(tokens) > 1 {
			newInput += " " + strings.Join(tokens[1:], " ")
		}
		return strings.Fields(newInput)
	}
	return tokens
}
