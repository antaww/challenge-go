package piscine

func verif(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	sRune := []rune(s)
	lettre := true
	for i := 0; i < len(s); i++ {
		if verif(sRune[i]) == true && lettre {
			if sRune[i] >= 'a' && sRune[i] <= 'z' {
				sRune[i] += 'A' - 'a'
			}
			lettre = false
		} else if sRune[i] >= 'A' && sRune[i] <= 'Z' {
			sRune[i] += 'a' - 'A'
		} else if verif(sRune[i]) == false {
			lettre = true
		}
	}
	return string(sRune)
}
