package piscine

func ListSize(l *List) int {
	counter := 1
	if l.Head == nil {
		return 0
	}
	iterator := l.Head
	for ; iterator.Next != nil; iterator = iterator.Next {
		counter++
	}
	return counter
}
