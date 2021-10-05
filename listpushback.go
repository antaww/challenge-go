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
	n := &NodeL{Data: data}
	head := l.Head
	tail := l.Tail
	if head == nil {
		head = n
	} else if head != nil {
		tail.Next = n.Next
		n = tail
	}
}
