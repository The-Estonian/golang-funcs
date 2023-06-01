package piscine

func AppendRange(min, max int) []int {
	if max < min {
		return []int{}
	}
	returnSlice := []int{}
	for i := min; i < max; i++ {
		returnSlice = append(returnSlice, i)
	}
	return returnSlice
}
