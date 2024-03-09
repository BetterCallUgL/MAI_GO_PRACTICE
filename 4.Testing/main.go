package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("hui")

	files, _ := os.ReadDir("/proc/self/fd")
	for _, file := range files {
		val, err := os.Readlink("/proc/self/fd/" + file.Name())
		if err == nil {
			fmt.Println(val)
		}
	}
	// fmt.Print(lol, "huita")
	file.Close()
}
