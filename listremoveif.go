package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	newListNode := &List{}
	looper := l.Head
	for ; looper != nil; looper = looper.Next {
		if looper.Data != data_ref {
			// ListPushBack(prevNode, looper.Data)
			newNode := &NodeL{Data: looper.Data}
			if newListNode.Head == nil {
				newListNode.Head = newNode
			} else {
				innerLooper := newListNode.Head
				for ; innerLooper.Next != nil; innerLooper = innerLooper.Next {
				}
				innerLooper.Next = newNode
			}
		}
	}
	*l = *newListNode
}
