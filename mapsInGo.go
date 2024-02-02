package piscine

func Map(f func(int) bool, a []int) []bool {
	returnSlice := []bool{}
	for _, v := range a {
		returnSlice = append(returnSlice, f(v))
	}
	return returnSlice
}
