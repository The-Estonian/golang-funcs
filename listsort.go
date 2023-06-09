package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	for firstIndex := l; firstIndex != nil; firstIndex = firstIndex.Next {
		for secondIndex := firstIndex.Next; secondIndex != nil; secondIndex = secondIndex.Next {
			if firstIndex.Data > secondIndex.Data {
				firstIndex.Data, secondIndex.Data = secondIndex.Data, firstIndex.Data
			}
		}
	}
	return l
}
