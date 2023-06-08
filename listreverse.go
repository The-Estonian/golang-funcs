package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	} else {
		reverseList := &List{}
		looper := l.Head
		for ; looper != nil; looper = looper.Next {
			ListPushFront(reverseList, looper.Data)
		}
		secondLooper := reverseList.Head
		for ; secondLooper != nil; secondLooper = secondLooper.Next {
			l.Tail = secondLooper
		}
		l.Head = reverseList.Head
	}
}
