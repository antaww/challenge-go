package piscine

func ConcatParams(args []string) string {
	phrase := ""
	j := 0
	for _, i := range args {
		phrase = phrase + i
		j = j + 1
		if j < len(args) {
			phrase += "\n"
		}
	}
	return phrase
}
