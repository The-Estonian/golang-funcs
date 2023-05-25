package piscine

func StrLen(s string) int {
	stringByte := []rune(s)
	counter := 0
	for i := 0; i < len(stringByte); i++ {
		counter++
	}
	return counter
}
