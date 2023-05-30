package piscine

func Capitalize(s string) string {
	stringRune := []rune(s)
	returnString := ""
	counter := 0
	for i := 0; i < len(stringRune); i++ {
		if IsAlpha(string(stringRune[i])) {
			counter++
		} else {
			counter = 0
		}

		if counter == 1 {
			if stringRune[i] >= 97 && stringRune[i] <= 122 {
				returnString += string(stringRune[i] - 32)
			} else {
				returnString += string(stringRune[i])
			}
		} else {
			if stringRune[i] <= 90 && stringRune[i] >= 65 {
				returnString += string(stringRune[i] + 32)
			} else {
				returnString += string(stringRune[i])
			}
		}
	}

	return returnString
}
