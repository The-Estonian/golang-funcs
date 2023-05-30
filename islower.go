package piscine

func IsLower(s string) bool {
	runeString := []rune(s)
	trigger := true
	for i := 0; i < len(runeString); i++ {
		if runeString[i] >= 97 && runeString[i] <= 122 {
			trigger = true
		} else {
			return false
		}
	}
	return trigger
}
