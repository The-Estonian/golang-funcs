package piscine

func ListPushFront(l *List, data interface{}) {
	input := &NodeL{Data: data}
	oldInput := l.Head
	l.Head = input
	l.Head.Next = oldInput
}
