package piscine

func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count += 1
		l.Head = l.Head.Next
	}
	return count
}
