package piscine

func NRune(s string, n int) rune {
	nb := n
	if len(s) > nb {
		return 0
	} else {
		return []rune(s)[n]
	}
}
