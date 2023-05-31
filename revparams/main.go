package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	argList := []string{}
	for i := 0; i < len(arguments); i++ {
		argList = append(argList, arguments[i])
	}
	for x := len(argList) - 1; x >= 0; x-- {
		for _, j := range argList[x] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
