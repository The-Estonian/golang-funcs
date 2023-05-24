package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for firstLoop := 0; firstLoop <= 9; firstLoop++ {
		for secondLoop := 0; secondLoop <= 9; secondLoop++ {
			fourthLoop := secondLoop + 1
			for thirdLoop := firstLoop; thirdLoop <= 9; thirdLoop++ {
				for ; fourthLoop <= 9; fourthLoop++ {
					z01.PrintRune(rune(firstLoop + '0'))
					z01.PrintRune(rune(secondLoop + '0'))
					z01.PrintRune(' ')
					z01.PrintRune(rune(thirdLoop + '0'))
					z01.PrintRune(rune(fourthLoop + '0'))
					if firstLoop < 9 || secondLoop < 8 || thirdLoop < 9 || fourthLoop < 9 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				fourthLoop = 0
			}
		}
	}
	z01.PrintRune('\n')
}
