package piscine

func LastRune(s string) rune {
	characters := []rune(s)
	return characters[len(characters)-1]
}
