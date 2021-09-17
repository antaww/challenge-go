package piscine

func Capitalize(s string) string {
	var phrase string
	for _, letter := range s {
		if (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z') || (letter >= '0' && letter <= '9') {
			if s[letter-1] == ' ' {
				phrase += string(letter - 32)
			} else {
				phrase += string(letter)
			}
		} else {
			phrase += string(letter)
		}
	}
	return phrase
}
