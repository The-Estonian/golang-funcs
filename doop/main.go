package main

import (
	"os"
)

func atoiConverter2(n string) int {
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

func itoa(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	} else if n == 2 {
		return "2"
	} else if n == 3 {
		return "3"
	} else if n == 4 {
		return "4"
	} else if n == 5 {
		return "5"
	} else if n == 6 {
		return "6"
	} else if n == 7 {
		return "7"
	} else if n == 8 {
		return "8"
	} else if n == 9 {
		return "9"
	} else {
		return "ERROR"
	}
}

func BasicAtoi2(s string) int {
	sum := 0
	stringTrigger := false
	negTrigger := false
	if string(s[0]) == "-" {
		negTrigger = true
		s = s[1:]
	}
	for i := 0; i < len(s); i++ {
		if atoiConverter2(string(s[i])) == -1 {
			stringTrigger = true
		}
		power := 1
		if i < len(s) {
			for j := 1; j < len(s)-i; j++ {
				power *= 10
			}
		}
		sum += atoiConverter2(string(s[i])) * power
	}
	if negTrigger {
		sum *= -1
	}
	if stringTrigger {
		return 0
	} else {
		return sum
	}
}

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return true
	}
	if string(s[0]) == "-" {
		s = s[1:]
	}
	trigger := false
	stringRune := []rune(s)
	for i := 0; i < len(stringRune); i++ {
		if stringRune[i] <= 57 && stringRune[i] >= 48 {
			trigger = true
		} else {
			return false
		}
	}
	return trigger
}

func ReturnNbr(n int) string {
	collectedNumbers := ""
	returnString := ""
	trigger := false
	if n < 0 {
		n *= -1
		trigger = true
	}
	for x := 0; x <= n; x++ {
		if n >= 10 {
			collectedNumbers += itoa(n % 10)
			n /= 10
			x--
		} else {
			collectedNumbers += itoa(n)
			break
		}
	}
	for y := len(collectedNumbers) - 1; y >= 0; y-- {
		if trigger {
			returnString += "-"
			trigger = false
		}
		returnString += string(collectedNumbers[y])
	}
	return returnString
}

func adding(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func modulo(a, b int) int {
	return a % b
}

func main() {
	args := os.Args[1:]
	if !IsNumeric(args[0]) || !IsNumeric(args[2]) {
		return
	}
	if BasicAtoi2(args[0]) > 9223372036854775807 ||
		BasicAtoi2(args[0]) < -9223372036854775807 ||
		BasicAtoi2(args[2]) > 9223372036854775807 ||
		BasicAtoi2(args[2]) < -9223372036854775807 {
		return
	}
	if args[1] == "+" {
		if adding(BasicAtoi2(args[0]), BasicAtoi2(args[2])) > 9223372036854775807 {
			return
		}
		os.Stdout.WriteString(ReturnNbr(adding(BasicAtoi2(args[0]), BasicAtoi2(args[2]))) + "\n")
	} else if args[1] == "-" {
		if subtract(BasicAtoi2(args[0]), BasicAtoi2(args[2])) < -9223372036854775807 {
			return
		}
		os.Stdout.WriteString(ReturnNbr(subtract(BasicAtoi2(args[0]), BasicAtoi2(args[2]))) + "\n")
	} else if args[1] == "*" {
		if multiply(BasicAtoi2(args[0]), BasicAtoi2(args[2])) < -9223372036854775807 ||
			multiply(BasicAtoi2(args[0]), BasicAtoi2(args[2])) > 9223372036854775807 {
			return
		}
		os.Stdout.WriteString(ReturnNbr(multiply(BasicAtoi2(args[0]), BasicAtoi2(args[2]))) + "\n")
	} else if args[1] == "/" {
		if divide(BasicAtoi2(args[0]), BasicAtoi2(args[2])) < -9223372036854775807 ||
			divide(BasicAtoi2(args[0]), BasicAtoi2(args[2])) > 9223372036854775807 {
			return
		}
		if BasicAtoi2(args[2]) == 0 {
			os.Stdout.WriteString("No division by 0" + "\n")
			return
		}
		os.Stdout.WriteString(ReturnNbr(divide(BasicAtoi2(args[0]), BasicAtoi2(args[2]))) + "\n")
	} else if args[1] == "%" {
		if BasicAtoi2(args[2]) == 0 {
			os.Stdout.WriteString("No modulo by 0" + "\n")
			return
		}
		os.Stdout.WriteString(ReturnNbr(modulo(BasicAtoi2(args[0]), BasicAtoi2(args[2]))) + "\n")
	}
}
