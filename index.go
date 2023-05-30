package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	if len(toFind) > len(s) {
		return -1
	}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == toFind {
			return i
		}
		if string(s[i]) == string(toFind[0]) {
			iteratorCounter := i + len(toFind) - 1
			counter := 0
			for x := i; x <= iteratorCounter; x++ {
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
