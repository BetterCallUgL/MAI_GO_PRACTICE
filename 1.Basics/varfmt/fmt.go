//go:build !solution

package varfmt

import (
	"fmt"
	"strings"
)

func Sprintf(format string, args ...interface{}) string {
	cnt1 := 0
	cnt2 := 0
	flag := false
	empty := true
	var builder strings.Builder
	for _, r := range format {
		if r == '{' {
			flag = true
		} else if r == '}' {
			flag = false
			if empty {
				builder.WriteString(fmt.Sprint(args[cnt2]))
			} else {
				builder.WriteString(fmt.Sprint(args[cnt1]))
			}
			cnt1 = 0
			cnt2++
			empty = true
		} else {
			if !flag {
				builder.WriteRune(r)
			} else {
				empty = false
				cnt1 = cnt1*10 + int(r-'0')
			}
		}
	}

	return builder.String()
}
