package piscine

func SplitWhiteSpaces(s string) []string {
	returnSlice := []string{}
	cutOff := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "\n" || string(s[i]) == " " || string(s[i]) == "	" {
			returnSlice = append(returnSlice, cutOff)
			cutOff = ""
		} else {
			cutOff += string(s[i])
		}
	}
	for i := 0; i < len(cutOff); i++ {
		if string(cutOff[i]) != " " {
			returnSlice = append(returnSlice, cutOff[i:])
		}
	}

	return returnSlice
}
