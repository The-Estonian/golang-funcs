package piscine

func Capitalize(s string) string {
	stringRune := []rune(s)
	returnString := ""
	if stringRune[0] >= 97 {
		returnString += string(stringRune[0] - 32)
	} else {
		returnString += string(stringRune[0])
	}
	for i := 1; i < len(stringRune); i++ {
		if stringRune[i-1] == 32 {
			if stringRune[i] >= 97 && stringRune[i] <= 122 {
				returnString += string(stringRune[i] - 32)
			} else {
				returnString += string(stringRune[i])
			}
		} else {
			if stringRune[i] <= 90 || stringRune[i] >= 65 {
				returnString += string(stringRune[i])
			} else {
				returnString += string(stringRune[i] + 32)
			}
		}
	}
	return returnString
}
