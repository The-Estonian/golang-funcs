package piscine

func Sqrt(nb int) int {
	num := 1
	for i := 1; i < 10; i++ {
		num = (num + nb/num) / 2
	}
	if num%2 == 0 {
		return num
	} else {
		return 0
	}
}
