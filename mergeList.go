package piscine

func ListMerge(l1 *List, l2 *List) {
	secondLooper := l2.Head
	for ; secondLooper != nil; secondLooper = secondLooper.Next {
		firstLooper := l1.Head
		insertData := &NodeL{Data: secondLooper.Data}
		if l1.Head == nil {
			l1.Head = insertData
		} else {
			for ; firstLooper.Next != nil; firstLooper = firstLooper.Next {
			}
			firstLooper.Next = insertData
		}
	}
}
