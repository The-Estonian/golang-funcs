package piscine

func IsPrintable(s string) bool {
	stringRune := []rune(s)
	for i := 0; i < len(stringRune); i++ {
		if stringRune[i] < 32 {
			return false
		}
	}
	return true
}
