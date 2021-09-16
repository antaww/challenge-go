package piscine

func ToUpper(s string) string {
	var phrase string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			phrase = string(letter - 32)
		}
		phrase += string(letter)
	}
	return phrase
}
