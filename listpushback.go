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
	n := &NodeL{}
	if l.Head == nil {
		l.Head = n
		return
	}
	head := l.Head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = n
}
