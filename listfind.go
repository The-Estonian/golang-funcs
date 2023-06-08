package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	looper := l.Head
	for ; looper != nil; looper = looper.Next {
		if comp(looper.Data, ref) {
			return &looper.Data
		}
	}
	return nil
}
