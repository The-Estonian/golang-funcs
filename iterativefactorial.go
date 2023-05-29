package piscine

func IterativeFactorial(nb int) int {
	sum := 1
	for i := 1; i <= nb; i++ {
		sum *= i
	}
	return sum
}
