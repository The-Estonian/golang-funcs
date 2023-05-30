package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return true
	}
	trigger := false
	stringRune := []rune(s)
	for i := 0; i < len(stringRune); i++ {
		if stringRune[i] <= 57 && stringRune[i] >= 48 {
			trigger = true
		} else {
			return false
		}
	}
	return trigger
}
