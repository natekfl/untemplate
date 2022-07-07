package untemplate

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// The error returned when Extract is not successful
var ErrNoMatch = errors.New("unable to match")

var tempToken = "__TMPTKN__"

var delimiters = []rune{'{', '}'}

// An Untemplater is a representation of how to deconstruct a given input string into tokens.
// Create it using untemplate.Create
type Untemplater struct {
	source       string
	tokens       []string
	patternRegex *regexp.Regexp
}

// Creates an Untemplater given a source containing token literals
func Create(source string) (*Untemplater, error) {
	tokenRegex := regexp.MustCompile(string(delimiters[0]) + "([^" + string(delimiters) + "\\t\\r\\n]+)" + string(delimiters[1]))
	tokenMatches := tokenRegex.FindAllStringSubmatch(source, -1)
	tokens := make([]string, len(tokenMatches))
	for i, v := range tokenMatches {
		tokens[i] = v[1] //Index 1 is the first (and only) submatch
	}

	substitutedTemplate := tokenRegex.ReplaceAllString(source, tempToken) // Substitution is required before escaping so that the capture group doesn't get escaped
	escapedSubstitutedTemplate := regexp.QuoteMeta(substitutedTemplate)
	escapedTemplate := strings.ReplaceAll(escapedSubstitutedTemplate, tempToken, "(.+)")
	patternRegex, err := regexp.Compile(escapedTemplate)
	if err != nil {
		return nil, fmt.Errorf("error when constructing regex: %v", err)
	}
	return &Untemplater{
		source:       source,
		tokens:       tokens,
		patternRegex: patternRegex,
	}, nil
}

// Provided an input and a template, returns a map of tokens to values
func (t *Untemplater) Extract(input string) (map[string]string, error) {

	matches := t.patternRegex.FindStringSubmatch(input)
	if len(matches) != len(t.tokens)+1 {
		return nil, ErrNoMatch
	}

	result := make(map[string]string)
	for i, t := range t.tokens {
		result[t] = matches[i+1]
	}

	return result, nil
}
