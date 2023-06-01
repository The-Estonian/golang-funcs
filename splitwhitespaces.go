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
	if len(cutOff) > 0 {
		if string(cutOff[0]) != " " {
			returnSlice = append(returnSlice, cutOff)
		} else {
			returnSlice = append(returnSlice, cutOff[1:])
		}
	}

	return returnSlice
}
