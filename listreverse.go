package piscine

func ListReverse(l *List) {
	actuel := l.Head
	precedent := l.Head
	precedent = nil
	for actuel != nil {
		next := actuel.Next
		actuel.Next = precedent
		precedent = actuel
		actuel = next
	}
	temp := l.Head
	l.Head = l.Tail
	l.Tail = temp
}
