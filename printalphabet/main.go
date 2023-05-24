package main

import "github.com/01-edu/z01"

func main() {
	alphabetAsString := "abcdefghijklmnopqrstuvwxyz"
	alphabetAsRune := []rune(alphabetAsString)
	for i := 0; i < len(alphabetAsRune); i++ {
		z01.PrintRune(alphabetAsRune[i])
	}
	z01.PrintRune('\n')
}
