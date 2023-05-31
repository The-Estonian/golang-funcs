package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0][1:]
	index := 0
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == '/' {
			index = i
		}
	}
	for j := index + 1; j < len(arguments); j++ {
		z01.PrintRune(rune(arguments[j]))
	}
	z01.PrintRune(rune('\n'))
}
