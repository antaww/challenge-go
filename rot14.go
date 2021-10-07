package piscine

func Rot14(s string) string {
	sRune := []rune(s)
	for i := 0; i < len(s); i++ {
		t := sRune[i]
		if sRune[i] >= 'a' && sRune[i] <= 'z' {
			if sRune[i]-12 < 'a' {
				sRune[i] = 'z' + ((t - 'a') - 12)
			}
			sRune[i] -= 12
		}
		if sRune[i] >= 'A' && sRune[i] <= 'Z' {
			if sRune[i]-12 < 'A' {
				sRune[i] = 'Z'
			}
			sRune[i] -= 12
		}
	}
	return string(sRune)
}
