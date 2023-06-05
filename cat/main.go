package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			text, err := ioutil.ReadFile(args[i])
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
				printStr("\n")
				os.Exit(1)
			}
			printStr(string(text))
			printStr("\n")
		}
	} else {
		x := 0
		for x < 10 {
			buffer := make([]byte, 1024)
			sentence, _ := os.Stdin.Read(buffer)
			printStr(string(buffer[:sentence]))
			x++
		}
	}
}
