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
	args := os.Args
	if len(args) > 1 {
		for i := 0; i < len(args); i++ {
			text, err := ioutil.ReadFile(args[i])
			if err != nil {
				printStr("ERROR: ")
				printStr(err.Error())
			}
			z01.PrintRune('\n')
			printStr(string(text))
		}
	} else {
	}
	for {
		buffer := make([]byte, 1024)
		ui, _ := os.Stdin.Read(buffer)
		printStr(string(buffer[:ui]))
	}
}
