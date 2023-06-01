package piscine

func AppendRange(min, max int) []int {
	if max < min || max == 0 && min == 0 {
		return []int(nil)
	}
	returnSlice := []int{}
	for i := min; i < max; i++ {
		returnSlice = append(returnSlice, i)
	}
	return returnSlice
}
