package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	trigger := false
	for firstLoop := 0; firstLoop <= 7; firstLoop++ {
		for secondLoop := firstLoop + 1; secondLoop <= 8; secondLoop++ {
			for thirdLoop := secondLoop + 1; thirdLoop <= 9; thirdLoop++ {
				if trigger {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				trigger = true
				z01.PrintRune(rune('0' + firstLoop))
				z01.PrintRune(rune('0' + secondLoop))
				z01.PrintRune(rune('0' + thirdLoop))
			}
		}
	}
	z01.PrintRune('\n')
}
