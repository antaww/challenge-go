package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if !(letter >= 'A' && letter <= 'Z') {
			return false
		}
	}
	return true
}
