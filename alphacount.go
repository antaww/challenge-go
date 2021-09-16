package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, letter := range s {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			counter++
		}
	}
	return counter
}
