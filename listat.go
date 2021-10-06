package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	if pos == 0 {
		return l
	}
	for count <= pos {
		if l == nil {
			return nil
		}
		l = l.Next
		count++
	}
	return l
}
