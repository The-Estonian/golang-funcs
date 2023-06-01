package piscine

func PrintNbrBaseReturn(nbr int, base string) string {
	numberSort := []int{}
	for i := 0; i <= nbr; i++ {
		if nbr >= len(base) {
			numberSort = append(numberSort, nbr%len(base))
			nbr /= len(base)
			i--
		} else if nbr < len(base) {
			numberSort = append(numberSort, nbr)
			break
		}
	}
	returnString := ""
	for i := len(numberSort) - 1; i >= 0; i-- {
		returnString += string(base[numberSort[i]])
	}
	return returnString
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	decimalNbr := AtoiBase(nbr, baseFrom)
	return PrintNbrBaseReturn(decimalNbr, baseTo)
}
