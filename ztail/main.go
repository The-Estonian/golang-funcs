package main

import (
	"fmt"
	"os"
)

func atoiConverter(n string) int {
	if n == "0" {
		return 0
	} else if n == "1" {
		return 1
	} else if n == "2" {
		return 2
	} else if n == "3" {
		return 3
	} else if n == "4" {
		return 4
	} else if n == "5" {
		return 5
	} else if n == "6" {
		return 6
	} else if n == "7" {
		return 7
	} else if n == "8" {
		return 8
	} else if n == "9" {
		return 9
	} else {
		return -1
	}
}

func main() {
	args := os.Args[1:]
	if args[0] == "-c" {
		number := args[1]
		args = args[2:]
		if len(args) > 0 {
			for i := 0; i < len(args); i++ {
				text, err := os.ReadFile(args[i])
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				if i > 0 {
					fmt.Println("")
				}
				fmt.Printf("==> %v <==\n", args[i])
				for j := len(text) - atoiConverter(number) + 1; j < len(text); j++ {
					fmt.Print(string(text[j]))
				}
				fmt.Println("")
			}
		}
	}
}
