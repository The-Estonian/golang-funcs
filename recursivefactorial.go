package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 25 {
		return 0
	}
	if nb == 1 {
		return nb
	}
	return nb * RecursiveFactorial(nb-1)
}
