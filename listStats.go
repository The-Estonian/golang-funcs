package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	if l.Data == nil {
		return nil
	} else {
		looper := l
		for ; looper != nil; looper = looper.Next {
			if counter == pos {
				return looper
			}
			counter++
		}
	}
	return nil
}
