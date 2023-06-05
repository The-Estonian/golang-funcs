package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) == 1 {
		text, _ := ioutil.ReadFile(args[0])
		fmt.Println(string(text))
	} else {
		fmt.Println("Too many arguments")
	}
}
