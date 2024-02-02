package piscine

func Sqrt(nb int) int {
	num := 0
	for i := 0; i <= nb; i++ {
		if i*i == nb {
			num = i
		}
	}
	return num
}
