package piscine

func Capitalize(s string) string {
	sRune := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if (sRune[i] >= 'A' && sRune[i] <= 'Z') || (sRune[i] >= 'a' && sRune[i] <= 'z') || (sRune[i] >= '0' && sRune[i] <= '9') {
			if s[i-1] == ' ' {
				sRune[i] += sRune[i-32]
			} else {
				sRune[i] += sRune[i]
			}
		} else {
			sRune[i] += sRune[i]
		}
	}
	return string(sRune)
}
