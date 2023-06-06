package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascending := true
	answer := true
	if a[0] > a[1] {
		ascending = false
	}
	for x := 0; x < len(a); x++ {
		for y := x + 1; y < len(a); y++ {
			if ascending {
				if f(a[x], a[y]) < 0 {
					answer = true
				} else if f(a[x], a[y]) == 0 {
					continue
				} else {
					return false
				}
			} else {
				if f(a[x], a[y]) > 0 {
					answer = true
				} else if f(a[x], a[y]) == 0 {
					continue
				} else {
					return false
				}
			}
		}
	}
	return answer
}
