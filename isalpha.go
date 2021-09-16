package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if !((letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9')) {
			return false
		}
	}
	return true
}
