package piscine

func ListSize(l *List) int {
	counter := 1
	iterator := l.Head
	for ; iterator.Next != nil; iterator = iterator.Next {
		counter++
	}
	return counter
}
