package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	input := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = input
	} else {
		loopVariable := l.Head
		for ; loopVariable.Next != nil; loopVariable = loopVariable.Next {
		}
		loopVariable.Next = input
	}
}
