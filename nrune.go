package piscine

func NRune(s string, n int) rune {
	nb := n
	if nb > len(s) {
		return 0
	} else {
		n := nb - 1
		return []rune(s)[n]
	}
}
