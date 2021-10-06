package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	count2 := 0
	for l != nil {
		count += 1
		l = l.Next
	}
	if pos <= count {
		for l != nil {
			if count2 == pos {
				return l
			}
			count2 += 1
			l = l.Next
		}
	}
	return nil
}
