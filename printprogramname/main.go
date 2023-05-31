package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0][2:]
	argumentsRunes := []rune(arguments)
	for i := 0; i < len(argumentsRunes); i++ {
		z01.PrintRune(rune(arguments[i]))
	}
	z01.PrintRune(rune('\n'))
}
