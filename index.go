package piscine

func Index(s string, toFind string) int {
	mot := []rune(s)
	toFindword := []rune(toFind)
	TailleMot := len(mot)
	TailleFind := len(toFindword)
	for i := 0; i <= TailleMot-TailleFind; i++ {
		if toFind == s[i:i+TailleFind] {
			return (i)
		}
	}
	return -1
}
