package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if !(letter == 92) {
			return false
		}
	}
	return true
}
