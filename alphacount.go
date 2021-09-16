package piscine

func AlphaCount(s string) int {
	counter := 0
	str := []rune(s)

	for i := 0; i <= len(str); i++ {
		if (str[i] >= 65 && str[i] <= 90) || (str[i] >= 97 && str[i] <= 122) {
			counter++
		}
	}
	return counter
}
