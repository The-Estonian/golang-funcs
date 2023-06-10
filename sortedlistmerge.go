package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	looper := n2
	for ; looper != nil; looper = looper.Next {
		SortListInsert(n1, looper.Data)
	}
	return n1
}
