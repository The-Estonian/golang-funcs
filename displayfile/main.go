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
		file, _ := os.Open(args[0])
		text, _ := ioutil.ReadAll(file)
		fmt.Println(string(text))
		defer file.Close()
	} else {
		fmt.Println("Too many arguments")
	}
}
