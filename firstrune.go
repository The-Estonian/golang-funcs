package piscine

func FirstRune(s string) rune {
	if rune(s[0]) < 65 || rune(s[0]) > 127 {
		return rune(s[0] - 223)
	}
	return rune(s[0])
}
