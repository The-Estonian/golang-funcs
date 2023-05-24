package main

import "github.com/01-edu/z01"

func main() {
	alphabetAsString := "abcdefghijklmnopqrstuvwxyz"
	alphabetAsRune := []rune(alphabetAsString)
	for i := len(alphabetAsRune); i > 0; i-- {
		z01.PrintRune(alphabetAsRune[i-1])
	}
	z01.PrintRune('\n')
}
