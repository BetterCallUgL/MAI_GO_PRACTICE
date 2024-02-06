//go:build !solution

package spacecollapse

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func CollapseSpaces(input string) string {
	var builder strings.Builder
	builder.Grow(len(input))
	var prev_rune rune

	for len(input) > 0 {
		r, size := utf8.DecodeRuneInString(input)
		if !unicode.IsSpace(r) {
			builder.WriteRune(r)
		} else {
			r = ' '
			if prev_rune != ' ' {
				builder.WriteRune(r)
			}
		}
		input = input[size:]
		prev_rune = r
	}

	return builder.String()
}
