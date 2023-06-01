package piscine

func ConcatParams(args []string) string {
	newString := ""
	for i := 0; i < len(args); i++ {
		newString += args[i]
		if i != len(args)-1 {
			newString += "\n"
		}
	}
	return newString
}
