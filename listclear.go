package piscine

func ListClear(l *List) {
	l.Head.Data = nil
	l.Tail.Data = nil
}
