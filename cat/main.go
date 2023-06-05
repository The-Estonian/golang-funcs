package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		for i := 0; i < len(args); i++ {
			text, err := ioutil.ReadFile(args[i])
			if err != nil {
				fmt.Printf("ERROR: %s\n", err)
			}
			fmt.Println(string(text))
		}
	} else {
		for {
			var input string
			fmt.Scanln(&input)
			fmt.Println(input)
		}
	}
}
