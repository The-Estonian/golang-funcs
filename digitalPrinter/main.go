package main

import "github.com/01-edu/z01"

func main() {
	alphabetAsString := "0123456789"
	alphabetAsRune := []rune(alphabetAsString)
	for i := 0; i < len(alphabetAsRune); i++ {
		z01.PrintRune(alphabetAsRune[i])
	}
	z01.PrintRune('\n')
}
