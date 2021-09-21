package piscine

func StrRev(s string) string {
	sRune := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		sRune[i], sRune[j] = sRune[j], sRune[i]
	}
	return string(s)
}
