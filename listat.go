package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for l.Next != nil {
		count += 1
		l = l.Next
	}
	if pos <= count {
		for l != nil {
			if count == pos {
				return l
			}
		}
		return nil
	}
	return nil
}
