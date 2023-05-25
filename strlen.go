package piscine

func StrLen(s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		counter++
	}
	return counter
}
