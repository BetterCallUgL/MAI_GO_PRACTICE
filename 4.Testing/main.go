package main

import (
	"fmt"
)

func assert(val interface{}) {
	if _, ok := val.(int); !ok {
		fmt.Print("lol")
	}
}

func main() {
	lol := int8(1)
	assert(lol)
}
