package piscine

func Index(s string, toFind string) int {
	if len(s) > 0 && len(toFind) > 0 && len(s) > len(toFind) {
		for i := 0; i < len(s); i++ {
			if string(s[i]) == toFind {
				return i
			}
			if string(s[i]) == string(toFind[0]) {
				counter := 0
				for x := i; x <= len(s); x++ {
					if string(s[x]) == string(toFind[x-i]) {
						counter++
					}
					if counter == len(toFind) {
						return i
					}
				}

			}
		}
	}
	return -1
}
