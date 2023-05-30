package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	toRun := true
	if len(base) >= 2 {
		for x := 0; x < len(base); x++ {
			for y := x + 1; y < len(base); y++ {
				if base[x] == base[y] || base[x] == '-' || base[x] == '+' {
					z01.PrintRune(rune('N'))
					z01.PrintRune('V')
					toRun = false
					break
				}
			}
		}
		if toRun {
			numberSort := []int{}
			trigger := false
			if nbr < 0 {
				nbr *= -1
				trigger = true
			}
			for i := 0; i <= nbr; i++ {
				if nbr >= len(base) {
					numberSort = append(numberSort, nbr%len(base))
					nbr /= len(base)
					i--
				} else if nbr < len(base) {
					numberSort = append(numberSort, nbr)
					break
				}
			}
			if trigger {
				z01.PrintRune('-')
			}
			for i := len(numberSort) - 1; i >= 0; i-- {
				z01.PrintRune(rune(base[numberSort[i]]))
			}
		}
	}
}
