package piscine

func NRune(s string, n int) rune {
	nb := n - 1
	if nb > len(s) {
		return 0
	} else {
		return []rune(s)[nb]
	}
}
