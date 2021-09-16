package piscine

func ToUpper(s string) string {
	var phrase string
	for i := 0; i <= len(s)-1; i++ {
		if i >= 'a' && i <= 'z' {
			phrase += string(s[i] - 32)
		}
		if i >= 'A' && i <= 'Z' {
			phrase += string(s[i])
		}
	}
	return phrase
}
