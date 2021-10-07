package piscine

func Rot14(s string) string {
	sRune := []rune(s)
	for i := 0; i < len(s); i++ {
		if verif(sRune[i]) == true {
			if sRune[i] >= 'a' && sRune[i] <= 'z' {
				sRune[i] -= 12
			}
			if sRune[i] >= 'A' && sRune[i] <= 'Z' {
				sRune[i] -= 12
			}
		} else if verif(sRune[i]) == false {
		}
	}
	return string(sRune)
}
