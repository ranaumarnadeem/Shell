package shell

import (
	"errors"
	"regexp"
	"strings"
)

var (
	tokenPattern = regexp.MustCompile(`"[^"]*"|'[^']*'|[^\s]+`)

	aliases = make(map[string]string)
)

func ParseInput(input string) ([]string, error) {
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

	// Alias expansion for first token
	if aliasCmd, ok := aliases[tokens[0]]; ok {
		// re-tokenize alias command + the rest
		expanded := aliasCmd + " " + strings.Join(tokens[1:], " ")
		toks, err := ParseInput(expanded)
		if err != nil {
			return nil, err
		}
		return toks, nil
	}

	return tokens, nil
}

var (
	ErrEmptyInput = errors.New("input is empty")
)
