package piscine

func Capitalize(s string) string {
	sRune := []rune(s)
	lettre := true
	for i := 0; i < len(s); i++ {
		if ((sRune[i] >= 'A' && sRune[i] <= 'Z') || (sRune[i] >= 'a' && sRune[i] <= 'z') || (sRune[i] >= '0' && sRune[i] <= '9')) && lettre == true {
			if sRune[i] >= 'a' && sRune[i] <= 'z' {
				sRune[i] += sRune[i-32]
				lettre = false
			} else if sRune[i] >= 'A' && sRune[i] <= 'Z' {
				sRune[i] += sRune[i+32]
			}
		} else {
			lettre = true
		}
	}
	return string(sRune)
}
