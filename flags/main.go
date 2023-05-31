package main

import (
	"fmt"
	"os"
)

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table); j++ {
			if table[i] < table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func main() {
	arguments := os.Args[1:]
	orderTrigger := false
	if len(arguments) < 1 || arguments[0] == "--help" || arguments[0] == "-h" {
		fmt.Println(`--insert
  -i
         This flag inserts the string into the string passed as argument.
--order
  -o
         This flag will behave like a boolean, if it is called it will order the argument.`)
		return
	}
	cleanArguments := []string{}
	insertString := ""
	for i := 0; i < len(arguments); i++ {
		if arguments[i] == "-o" || arguments[i] == "--order" {
			orderTrigger = true
		} else {
			cutArg := false
			for j := 0; j < len(arguments[i]); j++ {
				if len(arguments[i][j:]) > 3 {
					if arguments[i][j] == '-' && arguments[i][j+1] == 'i' && arguments[i][j+2] == '=' {
						insertString += arguments[i][j+3:]
						cutArg = true
					}
				}
				if len(arguments[i][j:]) > 9 {
					if arguments[i][j] == '-' && arguments[i][j+1] == '-' && arguments[i][j+2] == 'i' {
						insertString += arguments[i][j+9:]
						cutArg = true
					}
				}
			}
			if !cutArg {
				cleanArguments = append(cleanArguments, arguments[i])
				cutArg = false
			}
		}
	}
	if !orderTrigger {
		fmt.Println(cleanArguments[0] + insertString)
	} else {
		concatString := cleanArguments[0] + insertString
		newString := []rune(concatString)
		for x := 0; x < len(newString); x++ {
			for y := 0; y < len(newString); y++ {
				if newString[x] < newString[y] {
					newString[x], newString[y] = newString[y], newString[x]
				}
			}
		}
		fmt.Println(string(newString))
	}
}
