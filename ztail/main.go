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
			for i := 0; i < len(args); i++ {
				byteText, err := os.ReadFile(args[i])
				text := []rune(string(byteText))
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				if i > 0 {
					fmt.Println("")
				}
				fmt.Printf("==> %v <==\n", args[i])
				if len(text) <= BasicAtoi(number) {
					answer := []rune{}
					for y := 0; y < len(text)-1; y++ {
						answer = append(answer, text[y])
					}
					fmt.Println(" " + string(answer))
				} else {
					for j := len(text) - BasicAtoi(number) + 1; j < len(text); j++ {
						fmt.Print(string(text[j]))
					}
					fmt.Println("")
				}
			}
			os.Exit(1)
		}
	}
}
