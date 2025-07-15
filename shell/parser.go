package shell

import (
	"errors"
	"regexp"
	"strings"
)

// tokenPattern matches quoted or unquoted tokens
var tokenPattern = regexp.MustCompile(`"[^"]*"|'[^']*'|[^\s]+`)

// ParseInput splits the command line into tokens and expands aliases
func ParseInput(input string, aliases map[string]string) ([]string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil, nil
	}

	rawTokens := tokenPattern.FindAllString(input, -1)
	if len(rawTokens) == 0 {
		return nil, nil
	}

	tokens := make([]string, len(rawTokens))
	for i, tok := range rawTokens {
		if (strings.HasPrefix(tok, "\"") && strings.HasSuffix(tok, "\"")) ||
			(strings.HasPrefix(tok, "'") && strings.HasSuffix(tok, "'")) {
			tok = tok[1 : len(tok)-1]
		}
		tokens[i] = tok
	}

	// Expand alias if the first token matches
	if aliasCmd, ok := aliases[tokens[0]]; ok {
		expanded := aliasCmd + " " + strings.Join(tokens[1:], " ")
		return ParseInput(expanded, aliases) // Recursive with alias context
	}

	return tokens, nil
}

// ErrEmptyInput is returned when input is just whitespace
var ErrEmptyInput = errors.New("input is empty")
