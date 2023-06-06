package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for x := 0; x < len(a); x++ {
		for y := x + 1; y < len(a); y++ {
			if f(a[x], a[y]) >= 0 {
				return false
			}
		}
	}
	return true
}
