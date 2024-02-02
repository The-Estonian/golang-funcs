package piscine

func StrRev(s string) string {
	reversedString := ""
	for i := len(s); i > 0; i-- {
		reversedString += string(s[i-1])
	}
	return reversedString
}
