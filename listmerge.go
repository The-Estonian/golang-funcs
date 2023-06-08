package piscine

func ListMerge(l1 *List, l2 *List) {
	firstLooper := l1.Head
	secondLooper := l2.Head
	for ; secondLooper != nil; secondLooper = secondLooper.Next {
		insertData := &NodeL{Data: secondLooper.Data}
		for firstLooper.Next != nil {
			firstLooper = firstLooper.Next
		}
		firstLooper.Next = insertData
	}
}
