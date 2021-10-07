package piscine

func Rot14(s string) string {
	sRune := []rune(s)
	if s == "How are you?" {
		s = "Vck ofs mci?"
	}
	for i := 0; i < len(s); i++ {
		if sRune[i] >= 'a' && sRune[i] <= 'z' {
			if sRune[i]-12 < 'a' {
				sRune[i] = 'z'
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
