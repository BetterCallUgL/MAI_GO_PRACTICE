//go:build !solution

package reverse

import (
	"strings"
	"unicode/utf8"
)

func Reverse(input string) string {
	var build strings.Builder
	build.Grow(len(input))

	for len(input) > 0 {
		r, size := utf8.DecodeLastRuneInString(input)
		build.WriteRune(r)
		input = input[:len(input)-size]
	}

	return build.String()
}
