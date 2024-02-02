package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	trigger := false
	stringRune := []rune(s)
	for i := 0; i < len(stringRune); i++ {
		if stringRune[i] <= 122 && stringRune[i] >= 97 || stringRune[i] <= 90 && stringRune[i] >= 65 || stringRune[i] <= 57 && stringRune[i] >= 48 {
			trigger = true
		} else {
			return false
		}
	}
	return trigger
}
