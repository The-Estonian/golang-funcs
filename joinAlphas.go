package piscine

func BasicJoin(elems []string) string {
	returnString := ""
	for i := 0; i < len(elems); i++ {
		returnString += elems[i]
	}
	return returnString
}
