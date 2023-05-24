package piscine

import (
	"github.com/01-edu/z01"
)

func numberInput(n int) {
	a := '0'

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for b := 1; b <= n%10; b++ {
		a++
	}

	for b := -1; b >= n%10; b-- {
		a++
	}
	if n/10 != 0 {
		numberInput(n / 10)
	}

	z01.PrintRune(a)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	numberInput(n)
}
