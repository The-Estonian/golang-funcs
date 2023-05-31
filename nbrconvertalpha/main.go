package main

import (
	"os"

	"github.com/01-edu/z01"
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
	arguments := os.Args[1:]
	incomingMessage := []int{}
	decoder := map[int]int{
		1:  97,
		2:  98,
		3:  99,
		4:  100,
		5:  101,
		6:  102,
		7:  103,
		8:  104,
		9:  105,
		10: 106,
		11: 107,
		12: 108,
		13: 109,
		14: 110,
		15: 111,
		16: 112,
		17: 113,
		18: 114,
		19: 115,
		20: 116,
		21: 117,
		22: 118,
		23: 119,
		24: 120,
		25: 121,
		26: 122,
	}
	if string(arguments[0]) == "--upper" {
		for x, r := range decoder {
			decoder[x] = r - 32
		}
		arguments = arguments[1:]
	}
	for i := 0; i < len(arguments); i++ {
		incomingMessage = append(incomingMessage, BasicAtoi(arguments[i]))
	}
	for i := 0; i < len(incomingMessage); i++ {
		if incomingMessage[i] > 26 || incomingMessage[i] < 1 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune(rune(decoder[incomingMessage[i]]))
		}
	}
	z01.PrintRune('\n')
}
