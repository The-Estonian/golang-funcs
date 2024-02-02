package piscine

func removeZeroes(s string) string {
	if string(s[0]) != "0" {
		return s
	}
	return removeZeroes(string(s[1:]))
}

func TrimAtoi(s string) int {
	returnString := ""
	trigger := false
	returnNumber := 0
	stringRune := []rune(s)
	for i := 0; i < len(stringRune); i++ {
		if stringRune[i] <= 57 && stringRune[i] >= 48 {
			returnString += string(stringRune[i])
		}
		if stringRune[i] == 45 && len(returnString) == 0 {
			trigger = true
		}
	}
	if len(returnString) == 0 {
		return 0
	}
	returnString = removeZeroes(returnString)
	power := 1
	for i := 0; i < len(returnString)-1; i++ {
		power *= 10
	}
	for i := 0; i < len(returnString); i++ {
		returnNumber += atoiConverter3(string(returnString[i])) * power
		power /= 10
	}

	if trigger {
		returnNumber *= -1
	}
	return returnNumber
}
