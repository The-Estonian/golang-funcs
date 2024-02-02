package piscine

func AlphaCount(s string) int {
	stringConv := []rune(s)
	counter := 0
	for i := 0; i < len(stringConv); i++ {
		if stringConv[i] >= 65 && stringConv[i] <= 90 || stringConv[i] >= 97 && stringConv[i] <= 122 {
			counter++
		}
	}
	return counter
}
