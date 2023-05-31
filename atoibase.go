package piscine

func AtoiBase(s string, base string) int {
	for a := 0; a < len(s); a++ {
		for b := a + 1; b < len(base); b++ {
			if base[a] == base[b] || base[b] == '-' || base[b] == '+' {
				return 0
			}
		}
	}
	sum := 0
	indexList := []int{}
	for x := 0; x < len(s); x++ {
		for y := 0; y < len(base); y++ {
			if s[x] == base[y] {
				indexList = append(indexList, y)
			}
		}
	}
	baseNum := 1
	for i := len(indexList) - 1; i >= 0; i-- {
		sum += indexList[i] * baseNum
		baseNum *= len(base)
	}
	return sum
}
