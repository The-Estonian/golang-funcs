package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for x := 0; x < len(a); x++ {
		for y := x + 1; y < len(a); y++ {
			if f(a[x], a[y]) == 1 {
				a[x], a[y] = a[y], a[x]
			}
		}
	}
}
