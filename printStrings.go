package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	byteString := []byte(s)
	for i := 0; i < len(byteString); i++ {
		z01.PrintRune(rune(byteString[i]))
	}
}
