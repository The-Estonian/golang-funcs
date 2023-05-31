package main

import(
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arguments := os.Args[0][2:]
	for i := 0; i < len(arguments); i++ {
		z01.PrintRune(rune(arguments[i]))
	}
}
