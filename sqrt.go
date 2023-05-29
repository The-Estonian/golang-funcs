package piscine

func Sqrt(nb int) int {
	num := 1
	for i := 0; i < 10; i++ {
		num = (num + nb/num) / 2
	}
	if num%2 == 1 {
		return 0
	}
	return num
}
