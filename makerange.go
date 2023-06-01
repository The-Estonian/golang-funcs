package piscine

func MakeRange(min, max int) []int {
	if max < min || max == 0 && min == 0 {
		return []int(nil)
	}
	returnSlice := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		returnSlice[i] = min + i
	}
	return returnSlice
}
