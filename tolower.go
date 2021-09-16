package piscine

func ToLower(s string) string {
	var phrase string
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			phrase += string(letter + 32)
		} else {
			phrase += string(letter)
		}
	}
	return phrase
}
