package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(arguments); j++ {
			if arguments[i] < arguments[j] {
				arguments[i], arguments[j] = arguments[j], arguments[i]
			}
		}
	}
	for x := 0; x < len(arguments); x++ {
		z01.PrintRune(rune(arguments[x][0]))
		z01.PrintRune('\n')
	}
}
