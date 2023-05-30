package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	numberSort := []int{}
	for i := 0; i <= n; i++ {
		if n > 10 {
			numberSort = append(numberSort, n%10)
			n /= 10
		} else {
			numberSort = append(numberSort, n)
			break
		}
	}
	SortIntegerTable(numberSort)

	for i := 0; i < len(numberSort); i++ {
		z01.PrintRune(rune(numberSort[i]))
	}
}
