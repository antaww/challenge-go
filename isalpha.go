package piscine

func IsAlpha(s string) bool {
	for _, letter := range s {
		if !(letter >= 'A' && letter <= 'Z' && letter >= 'a' && letter <= 'z') {
			return false
		}
	}
	return true
}
