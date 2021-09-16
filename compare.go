package piscine

func FirstRune(s string) rune {
	return []rune(s)[0]
}

func LastRune(s string) rune {
	return []rune(s)[len(s)-1]
}

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if FirstRune(a) != FirstRune(b) {
		return -1
	} else if LastRune(a) != LastRune(b) {
		return 1
	}
}
