package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == toFind {
			return i
		}
		if string(s[i]) == string(toFind[0]) {
			counter := 0
			for x := i; x < len(toFind)+1; x++ {
				if string(s[x]) == string(toFind[x-i]) {
					counter++
				}
			}
			if counter == len(toFind) {
				return i
			}
		}
	}
	return -1
}
