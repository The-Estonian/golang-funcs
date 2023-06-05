package main

import (
	"github.com/01-edu/z01"
)

func intString(n int) string {
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
		return "-1"
	}
}

func intToString(n int) string {
	returnString := ""
	for x := 0; x < n; x++ {
		if n > 10 {
			returnString = intString(n%10) + returnString
			n /= 10
		} else {
			returnString = intString(n) + returnString
			break
		}
	}
	return returnString
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	answerString := "x = " + intToString(points.x) + ", y = " + intToString(points.y)
	printStr(answerString)
}
