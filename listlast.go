package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		iterator := l.Head
		for ; iterator.Next != nil; iterator = iterator.Next {
		}
		return iterator.Data
	}
}
