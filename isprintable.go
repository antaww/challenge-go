package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if !(letter >= 0 && letter <= 32) {
			return false
		}
	}
	return true
}
