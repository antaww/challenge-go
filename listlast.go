package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return *l.Head
	} else {
		return *l.Tail
	}
}
