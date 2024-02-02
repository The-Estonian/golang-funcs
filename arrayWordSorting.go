package piscine

func SortWordArr(a []string) {
	for x := 0; x < len(a); x++ {
		for y := x + 1; y < len(a); y++ {
			if a[x] > a[y] {
				a[x], a[y] = a[y], a[x]
			}
		}
	}
}
