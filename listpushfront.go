package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = l.Head
		return
	}
	n.Next = l.Head
	l.Head = n
}
