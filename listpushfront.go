package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = n
		return
	}
	tail := l.Tail
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = n
}
