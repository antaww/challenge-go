package piscine

func PrintNbrInOrder(n int) {
	nbr := n
	var array []int{n}
	for nbr, n := range array {
		if nbr >= '0' && nbr <= '9' {
			if len(nbr)-2 > len(nbr)-1 {
				nbr += []rune(n - 1)
			} else {
				nbr += n
			}

		}

	}
	return nbr
}
