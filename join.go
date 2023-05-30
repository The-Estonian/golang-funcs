package piscine

func Join(strs []string, sep string) string {
	returnString := ""
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			returnString += strs[i] + sep
		} else {
			returnString += strs[i]
		}
	}
	return returnString
}
