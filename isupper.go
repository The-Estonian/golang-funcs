package piscine

func IsUpper(s string) bool {
	runeString := []rune(s)
	trigger := true
	for i := 0; i < len(runeString); i++ {
		if runeString[i] >= 65 && runeString[i] <= 90 {
			trigger = true
		} else {
			return false
		}
	}
	return trigger
}
