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

func BasicAtoi(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		power := 1
		if i < len(s) {
			for j := 1; j < len(s)-i; j++ {
				power *= 10
			}
		}
		sum += atoiConverter(string(s[i])) * power
	}
	return sum
}

func main() {
	args := os.Args[1:]
	if args[0] == "-c" {
		number := args[1]
		args = args[2:]
		if len(args) > 0 {
			counter := 0
			for _, i := range args {
				text, err := os.ReadFile(i)
				if err != nil {
					fmt.Println(err.Error())
					counter++
					continue
				}
				if counter > 0 {
					fmt.Println(``)
				}
				fmt.Printf("==> %v <==\n", i)
				if len(text) < BasicAtoi(number) {
					if text[len(text)-1] == '\n' {
						fmt.Printf("%s\n", text[:len(text)-1])
					} else {
						fmt.Printf("%s\n", text)
					}
					os.Exit(1)
				} else {
					fmt.Printf("%s\n", text[len(text)-BasicAtoi(number):])
					os.Exit(1)
				}
				counter++
			}
		}
	}
}
