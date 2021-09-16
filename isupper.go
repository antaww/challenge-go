package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			return false
		} else {
			return true
		}
	}
	return true
}
