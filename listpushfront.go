package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = l.Head
		return
	} else {
		m := &NodeL{Data: data}
		m.Next = l.Head
		l.Head = m
	}
}
