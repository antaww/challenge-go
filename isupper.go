package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			return false
		}
	}
	return true
}
