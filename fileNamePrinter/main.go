package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[0][1:]
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == '/' {
			arguments = arguments[i+1:]
		}
	}
	for _, letter := range arguments {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
