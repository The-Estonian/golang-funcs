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
	for x := 0; x < len(argList); x++ {
		for _, j := range argList[x] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
