package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	x := 1
	for i := 0; i < 10; i++ {
		x = (x + nb/x) / 2
	}
	if x%2 == 1 {
		return 0
	}
	return x
}
