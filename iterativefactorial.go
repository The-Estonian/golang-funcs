package piscine

func IterativeFactorial(nb int) int {
	if nb < 1 || nb > 25 {
		return 0
	}
	sum := 1
	for i := nb; i > 0; i-- {
		sum *= i
	}
	return sum
}
