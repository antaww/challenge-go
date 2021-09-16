package piscine

func NRune(s string, n int) rune {
	nb := n
	if nb > len(s) {
		return 0
	} else {
		return []rune(s)[int(n)-1]
	}
}
