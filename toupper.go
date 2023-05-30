package piscine

func ToUpper(s string) string {
	stringRune := []rune(s)
	newString := ""
	for i := 0; i < len(stringRune); i++ {
		if stringRune[i] <= 90 && stringRune[i] >= 65 {
			newString += string(stringRune[i])
		} else if stringRune[i] <= 122 && stringRune[i] >= 97 {
			newString += string(stringRune[i] - 32)
		} else {
			newString += string(stringRune[i])
		}
	}
	return newString
}
