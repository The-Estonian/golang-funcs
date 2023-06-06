package main

import (
	"os"
)

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
		PrintNbr(adding(BasicAtoi2(args[0]), BasicAtoi2(args[2])))
	} else if args[1] == "-" {
		if subtract(BasicAtoi2(args[0]), BasicAtoi2(args[2])) < -9223372036854775807 {
			return
		}
		PrintNbr(subtract(BasicAtoi2(args[0]), BasicAtoi2(args[2])))
	} else if args[1] == "main.go" {
		if multiply(BasicAtoi2(args[0]), BasicAtoi2(args[2])) < -9223372036854775807 ||
			multiply(BasicAtoi2(args[0]), BasicAtoi2(args[2])) > 9223372036854775807 {
			return
		}
		PrintNbr(multiply(BasicAtoi2(args[0]), BasicAtoi2(args[2])))
	} else if args[1] == "/" {
		if divide(BasicAtoi2(args[0]), BasicAtoi2(args[2])) < -9223372036854775807 ||
			divide(BasicAtoi2(args[0]), BasicAtoi2(args[2])) > 9223372036854775807 {
			return
		}
		if BasicAtoi2(args[2]) == 0 {
			PrintStr("No division by 0")
			return
		}
		PrintNbr(divide(BasicAtoi2(args[0]), BasicAtoi2(args[2])))
	} else if args[1] == "%" {
		if BasicAtoi2(args[2]) == 0 {
			PrintStr("No modulo by 0")
			return
		}
		PrintNbr(modulo(BasicAtoi2(args[0]), BasicAtoi2(args[2])))
	}
}
