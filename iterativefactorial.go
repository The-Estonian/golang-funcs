package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}
	sum := 1
	for i := 1; i <= nb; i++ {
		sum *= i
	}
	return sum
}
