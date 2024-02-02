package piscine

func Any(f func(string) bool, a []string) bool {
	trigger := false
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			trigger = true
		}
	}
	return trigger
}
