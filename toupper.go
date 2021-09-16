package piscine

func ToUpper(s string) string {
	var phrase string
	for i := 0; i <= len(phrase)-1; i++ {
		if i >= 'a' && i <= 'z' {
			phrase += string(s[i] - 32)
		}
	}
	return phrase
}
