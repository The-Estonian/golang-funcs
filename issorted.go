package piscine

func fFunction(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for x := 0; x < len(a); x++ {
		for y := x + 1; y < len(a); y++ {
			if f(a[x], a[y]) == 1 {
				return false
			}
		}
	}
	return true
}
