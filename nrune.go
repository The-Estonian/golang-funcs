package piscine

func NRune(s string, n int) rune {
	characters := []rune(s)
	for i := 0; i < len(characters); i++ {
		if i+1 == n {
			return characters[i]
		}
	}
	return 0
}
