package piscine

func IsNumeric(s string) bool {
	for _, letter := range s {
		if !(letter >= '0' && letter <= '9') {
			return false
		}
	}
	return true
}
