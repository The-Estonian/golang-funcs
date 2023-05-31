package main

import (
	"os"

	"github.com/01-edu/z01"
)

func checkIfVowel(n string) bool {
	if n == "a" || n == "A" || n == "e" || n == "E" || n == "o" || n == "O" || n == "u" || n == "U" || n == "i" || n == "I" {
		return true
	}
	return false
}

func main() {
	arguments := os.Args[1:]
	listOfVowels := []rune{}
	returnString := ""

	for i := 0; i < len(arguments); i++ {
		for j := 0; j < len(arguments[i]); j++ {
			if checkIfVowel(string(arguments[i][j])) {
				listOfVowels = append(listOfVowels, rune(arguments[i][j]))
			}
		}
		if i == len(arguments)-1 {
			returnString += arguments[i]
		} else {
			returnString += arguments[i] + " "
		}
	}
	counter := 0
	for i := 0; i < len(returnString); i++ {
		if checkIfVowel(string(returnString[i])) {
			z01.PrintRune(listOfVowels[len(listOfVowels)-1-counter])
			counter++
		} else {
			z01.PrintRune(rune(returnString[i]))
		}
	}
	z01.PrintRune('\n')
}
