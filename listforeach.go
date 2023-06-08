package piscine

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}

func ListForEach(l *List, f func(*NodeL)) {
	looper := l.Head
	for ; looper != nil; looper = looper.Next {
		f(looper)
	}
}
