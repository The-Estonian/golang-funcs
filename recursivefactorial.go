package piscine

func recursivefactorial(nb int) int {
	if nb == 1 {
		return nb
	}
	return nb * IterativeFactorial(nb-1)
}
