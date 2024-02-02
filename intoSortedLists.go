package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	input := &NodeI{Data: data_ref}
	nodeLooper := l
	if l == nil {
		l = input
	} else {
		if nodeLooper.Data > input.Data {
			input.Next = nodeLooper
			nodeLooper = input
			return nodeLooper
		}
		for ; nodeLooper.Data < data_ref; nodeLooper = nodeLooper.Next {
			if nodeLooper.Next != nil {
				if nodeLooper.Next.Data >= data_ref {
					input.Next = nodeLooper.Next
					nodeLooper.Next = input
				}
			} else {
				nodeLooper.Next = input
			}
		}
	}
	return l
}
