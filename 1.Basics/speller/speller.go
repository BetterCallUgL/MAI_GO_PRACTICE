//go:build !solution

package speller

import (
	"fmt"
)

var dict = map[int64]string{
	0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five",
	6: "six", 7: "seven", 8: "eight", 9: "nine", 10: "ten", 90: "ninety",
	11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen",
	16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty",
	30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy", 80: "eighty",
	100: "hundred", 1000: "thousand", 1000000: "million", 1000000000: "billion",
}

func Spell(n int64) string {
	if n < 0 {
		return "minus " + Spell(-n)
	}
	if n == 0 {
		return "zero"
	}

	copy := n
	var rank int64 = 1
	for copy >= 1000 {
		copy /= 1000
		rank *= 1000
	}

	line := ""
	if copy > 99 {
		line = dict[copy/100] + " " + dict[100]
	}

	if len(line) > 0 && line[len(line)-1] != ' ' && copy%100 > 0 {
		line += " "
	}
	copy = copy % 100
	if copy > 0 && copy < 20 {
		line += dict[copy]
	} else if copy > 19 {
		if copy%10 != 0 {
			line += dict[(copy/10)*10] + "-" + dict[copy%10]
		} else {
			line += dict[copy]
		}
	}

	if n >= 1000 {
		if n%rank != 0 {
			return line + " " + dict[rank] + " " + Spell(n%rank)
		} else {
			return line + " " + dict[rank]
		}
	} else {
		return line
	}
}

func main() {
	var num int64
	fmt.Scan(&num)

	fmt.Println(Spell(num))
}
