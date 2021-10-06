package piscine

func ListSize(l *List, data interface{}) {
	count := 0
	for l.Head.Next != nil {
		count += 1
	}
}
