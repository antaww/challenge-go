package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	count := 0
	for count != pos {
		count += 1
	}
	return l
}
