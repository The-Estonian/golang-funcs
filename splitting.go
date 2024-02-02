package piscine

func Split(s, sep string) []string {
	returnSlice := []string{}
	stringSlice := ""
	counter := 0
	for i := 0; i < len(s); i++ {
		if counter == 0 {
			if len(s[i:]) >= len(sep) {
				if string(s[i:len(sep)+i]) == sep {
					counter = len(sep)
					returnSlice = append(returnSlice, stringSlice)
					stringSlice = ""
				} else {
					stringSlice += string(s[i])
				}
			} else {
				stringSlice += string(s[i])
			}
		}
		if counter > 0 {
			counter--
		}
		if i == len(s)-1 {
			returnSlice = append(returnSlice, stringSlice)
		}
	}
	return returnSlice
}
